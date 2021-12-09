package layout

import (
	"fmt"
	"html/template"
	"log"
	"time"
)

var templates *template.Template

// Expected directory structure: templates/layouts and templates/components.
func init() {
	init_layouts()
	init_components()
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

type ILayout interface {
	Template() *template.Template
	Context() map[string]string
}

type Layout struct {
	template *template.Template
	context  map[string]string
}

func PrintTemplates() {
	log.Default().Println(templates.DefinedTemplates())
}

func AddTemplateFromFilepath(filepath string) {
	var err error
	templates, err = templates.ParseFiles(filepath)
	if err != nil {
		log.Fatal(err)
	}
}

func AddTemplateFromString(inline string) {
	name := fmt.Sprintf("%d", time.Now().Unix())
	content := fmt.Sprintf(`{{define "content"}}%s{{end}}`, inline)

	var err error
	templates, err = templates.New(name).Parse(content)
	if err != nil {
		log.Fatal(err)
	}
}

func NewLayout(template_name string, ctx map[string]string) *Layout {
	return &Layout{
		template: templates.Lookup(template_name),
		context:  ctx,
	}
}

func (l *Layout) Template() *template.Template {
	return l.template
}

func (l *Layout) Context() map[string]string {
	return l.context
}
