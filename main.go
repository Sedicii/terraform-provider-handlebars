package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sedicii/terraform-provider-handlebars/handlebars_template"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: handlebars_template.Provider})
}
