package component

import (
	"html/template"

	pkgLayout "github.com/michaeltelford/swoop/layout"
)

type (
	IComponent interface {
		Name() string
		Context() map[string]interface{}
		Content() template.HTML
	}

	Component struct {
		name         string
		context      map[string]interface{}
		templateName string
	}
)

func (c *Component) Name() string {
	return c.name
}

func (c *Component) Context() map[string]interface{} {
	return c.context
}

func (c *Component) TemplateName() string {
	return c.templateName
}

func (c *Component) Template() *template.Template {
	return pkgLayout.LookupTemplate(c.templateName)
}

func (c *Component) Content() template.HTML {
	return pkgLayout.ParseTemplate(c.Template(), c.Context())
}

func NewComponent(name, template string) *Component {
	posts := map[string]interface{}{
		"posts": []struct{ Title string }{{"My first post yay :-)"}},
	}

	return &Component{
		name:         name,
		templateName: template,
		context:      posts,
	}
}
