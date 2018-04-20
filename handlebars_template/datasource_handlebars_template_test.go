package handlebars_template

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/lumasepa/raymond"
)

var testProviders = map[string]terraform.ResourceProvider{
	"handlebars": Provider(),
}

func TestTemplateRendering(t *testing.T) {
	var cases = []struct {
		vars     string
		template string
		want     string
	}{
		{`{}`, `ABC`, `ABC`},
		{`{\"a\": \"foo\"}`, `{{a}}`, `foo`},
		{`{\"a\": true}`, `{{a}}`, `true`},
		{`{\"a\": false}`, `{{a}}`, `false`},
		{`{\"a\": 43}`, `{{a}}`, `43`},
		{`{\"a\": 43.1}`, `{{a}}`, `43.1`},
		{`{\"a\": {\"a\": 1, \"b\": 2}}`, `{{#each a}}{{@key}} = {{this}} {{/each}}`, `a = 1 b = 2 `},
		{`{\"a\": [\"h\", \"i\"]}`, `{{#each a}}{{this}}{{/each}}`, `hi`},
		{`{}`, `{{foo}}`, ``},
		{`{}`, `/`, `/`},
	}

	for _, tt := range cases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				r.TestStep{
					Config: testTemplateConfig(tt.template, tt.vars),
					Check: func(s *terraform.State) error {
						got := s.RootModule().Outputs["rendered"]
						if tt.want != got.Value {
							return fmt.Errorf("handlebars_template:\n%s\nvars:\n%s\ngot:\n%s\nwant:\n%s\n", tt.template, tt.vars, got, tt.want)
						}
						return nil
					},
				},
			},
		})
	}
}

func testTemplateConfig(template, vars string) string {
	return fmt.Sprintf(`
		data "handlebars_template" "t0" {
			template = "%s"
			json_context = "%s"
		}
		output "rendered" {
				value = "${data.handlebars_template.t0.rendered}"
		}`, template, vars)
}

func TestTfIf(t *testing.T) {
	source := `{{#if yep}}YEP !{{else}}NOP !{{/if}}`

	raymondTfExtraHelpers()

	testCases := []struct {
		in          string
		expectedOut string
	}{
		{"0", "NOP !"},
		{"1", "YEP !"},
		{"10", "YEP !"},
		{"text", "YEP !"},
		{"", "NOP !"},
	}

	for _, testCase := range testCases {
		ctx := map[string]interface{}{"yep": testCase.in}
		res, err := raymond.Render(source, ctx, nil)

		if err != nil {
			t.Error("tfif error", err)
		}

		if res != testCase.expectedOut {
			t.Errorf("Test error in case %s \n Expected: %s \n Got: %s", testCase.in, testCase.expectedOut, res)
		}
	}
}
