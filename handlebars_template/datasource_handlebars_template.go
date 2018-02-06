package handlebars_template

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/aymerick/raymond"
	"github.com/hashicorp/terraform/helper/schema"
)

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

	rendered, err := raymond.Render(template, jsonContext)

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
