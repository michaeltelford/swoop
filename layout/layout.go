package layout

import (
	"html/template"
	"log"
)

var (
	templates *template.Template
)

type (
	ILayout interface {
		Template() *template.Template
		Context() map[string]string
	}

	Layout struct {
		template *template.Template
		context  map[string]string
	}
)

// Expected directory structure: templates/layouts and templates/components.
func init() {
	init_layouts()
	init_components()
	addInlineTemplate()
}

func init_layouts() {
	var err error
	templates, err = templates.ParseGlob("templates/layouts/*")
	if err != nil {
		log.Fatal(err)
	}
}

func init_components() {
	var err error
	templates, err = templates.ParseGlob("templates/components/*")
	if err != nil {
		log.Fatal(err)
	}
}

func addInlineTemplate() {
	name := "inline_content"
	content := `{{define "content"}}{{.inline_content}}{{end}}`

	var err error
	templates, err = templates.New(name).Parse(content)
	if err != nil {
		log.Fatal(err)
	}
}

func (l *Layout) Template() *template.Template {
	return l.template
}

func (l *Layout) Context() map[string]string {
	return l.context
}

func NewLayout(template_name string, ctx map[string]string) *Layout {
	return &Layout{
		template: templates.Lookup(template_name),
		context:  ctx,
	}
}

func PrintTemplates() {
	log.Default().Println(templates.DefinedTemplates())
}
