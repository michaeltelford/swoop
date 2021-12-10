package layout

import (
	"bytes"
	"html/template"
	"log"
)

var (
	templates *template.Template
)

type (
	ILayout interface {
		Template() *template.Template
		Context() map[string]interface{}
	}

	Layout struct {
		template *template.Template
		context  map[string]interface{}
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
	content := `{{.content}}`

	var err error
	templates, err = templates.New(name).Parse(content)
	if err != nil {
		log.Fatal(err)
	}
}

func (l *Layout) Template() *template.Template {
	return l.template
}

func (l *Layout) Context() map[string]interface{} {
	return l.context
}

func NewLayout(template_name string, ctx map[string]interface{}) *Layout {
	return &Layout{
		template: templates.Lookup(template_name),
		context:  ctx,
	}
}

func NewLayoutFromLayout(layout *Layout, ctx map[string]interface{}) *Layout {
	merged_ctx := MergeContexts(layout.Context(), ctx)
	return NewLayout(layout.Template().Name(), merged_ctx)
}

func PrintTemplates() {
	log.Default().Println(templates.DefinedTemplates())
}

func LookupTemplate(name string) *template.Template {
	return templates.Lookup(name)
}

func ParseTemplate(tmpl *template.Template, ctx map[string]interface{}) template.HTML {
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, ctx); err != nil {
		log.Fatal(err)
	}
	return template.HTML(buffer.String())
}

// Merge two maps, overwriting values in the first map with values from the second.
func MergeContexts(ctx1, ctx2 map[string]interface{}) map[string]interface{} {
	new_ctx := make(map[string]interface{})

	for k, v := range ctx1 {
		new_ctx[k] = v
	}
	for k, v := range ctx2 {
		new_ctx[k] = v
	}

	return new_ctx
}
