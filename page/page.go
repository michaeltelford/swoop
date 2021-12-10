package page

import (
	"html/template"
	"log"

	pkgComponent "github.com/michaeltelford/swoop/component"
	pkgLayout "github.com/michaeltelford/swoop/layout"
)

type (
	IPage interface {
		Method() string            // GET, POST, etc.
		Route() string             // The page URL path e.g. "/foo/bar"
		Layout() *pkgLayout.Layout // The layout to use for the page
		Content() template.HTML    // The content of the page
	}

	Page struct {
		method        string
		route         string
		layout        *pkgLayout.Layout
		inlineContent template.HTML
		components    []pkgComponent.IComponent
	}
)

func (p *Page) Method() string {
	return p.method
}

func (p *Page) Route() string {
	return p.route
}

func (p *Page) Layout() *pkgLayout.Layout {
	return p.layout
}

func (p *Page) Content() template.HTML {
	ctx := p.layout.Context()

	if p.inlineContent != "" {
		ctx["content"] = p.inlineContent
	} else if len(p.components) > 0 {
		var componentContent template.HTML
		for _, component := range p.components {
			componentContent += component.Content()
		}
		ctx["content"] = componentContent
	} else {
		log.Fatal("No content or components to render")
	}

	return pkgLayout.ParseTemplate(p.layout.Template(), ctx)
}

func NewPageFromString(method, route string, layout *pkgLayout.Layout, content template.HTML) *Page {
	return &Page{
		method:        method,
		route:         route,
		layout:        layout,
		inlineContent: content,
	}
}

func NewPageFromComponents(method, route string, layout *pkgLayout.Layout, components []pkgComponent.IComponent) *Page {
	return &Page{
		method:     method,
		route:      route,
		layout:     layout,
		components: components,
	}
}
