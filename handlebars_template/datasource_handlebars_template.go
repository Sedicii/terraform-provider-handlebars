package handlebars_template

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/lumasepa/raymond"
	"strconv"
)

type NullEscaper struct{}

func (self NullEscaper) Escape(s string) string {
	return s
}

var templateOpts = raymond.TemplateOptions{
	Escaper: NullEscaper{},
}

func handlebarsTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRenderHandlebarsTemplate,

		Schema: map[string]*schema.Schema{
			"template": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Contents of the template",
			},
			"json_context": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "variables to substitute as a json",
			},
			"rendered": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "rendered template",
			},
		},
	}
}

func dataSourceRenderHandlebarsTemplate(d *schema.ResourceData, meta interface{}) error {
	template := d.Get("template").(string)
	jsonContextStr := d.Get("json_context").(string)

	rendered, err := renderHandlebarsTemplate(template, jsonContextStr)
	if err != nil {
		return err
	}
	d.Set("rendered", rendered)
	d.SetId(hash(rendered))
	return nil
}

type templateRenderError error

func renderHandlebarsTemplate(template string, jsonContextStr string) (string, error) {
	jsonContext := make(map[string]interface{})
	json.Unmarshal([]byte(jsonContextStr), &jsonContext)

	rendered, err := raymond.Render(template, jsonContext, &templateOpts)

	if err != nil {
		return "", templateRenderError(
			fmt.Errorf("failed to render handlebars_template: %v", err),
		)
	}

	return rendered, nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

func raymondTfExtraHelpers() {
	// If for terraform evals "1" to true and others to false
	raymond.RemoveHelper("if")

	coerceExpr := func(expr string) bool {
		i, err := strconv.Atoi(expr)
		if err == nil { // is a number eval as x > 0
			if i > 0 {
				return true
			} else {
				return false
			}
		} else if len(expr) > 0 { // is a string eval as notEmpty(str)
			return true
		}
		return false
	}

	raymond.RegisterHelper("if", func(conditional string, options *raymond.Options) string {
		if coerceExpr(conditional) {
			return options.Fn()
		}
		return options.Inverse()
	})

}

func init() {
	raymondTfExtraHelpers()
}
