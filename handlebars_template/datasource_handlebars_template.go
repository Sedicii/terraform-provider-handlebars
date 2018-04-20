package handlebars_template

import (
	"crypto/sha256"
	"encoding/hex"
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
			"context": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "variables to substitute",
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
	context := d.Get("context").(map[string]interface{})
	fmt.Sprintf("%+v", context)

	rendered, err := renderHandlebarsTemplate(template, context)
	if err != nil {
		return err
	}
	d.Set("rendered", rendered)
	d.SetId(hash(rendered))
	return nil
}

type templateRenderError error

func renderHandlebarsTemplate(template string, context map[string]interface{}) (string, error) {
	//x, err := json.Marshal(context)
	//if err != nil {
	//	return "", err
	//}
	//jsonContext := make(map[string]interface{})
	//err = json.Unmarshal(x, &jsonContext)
	//if err != nil {
	//	return "", err
	//}
	rendered, err := raymond.Render(template, context, &templateOpts)

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
