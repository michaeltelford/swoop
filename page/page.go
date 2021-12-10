package page

import (
	"html/template"
	"log"

	pkgComponent "github.com/michaeltelford/swoop/component"
	pkgLayout "github.com/michaeltelford/swoop/layout"
)

type (
	IPage interface {
		Route() string             // The page URL path e.g. "/foo/bar"
		Layout() *pkgLayout.Layout // The layout to use for the page
		Content() template.HTML    // The content of the page
	}

	Page struct {
		route          string
		layout         *pkgLayout.Layout
		inline_content string
		components     []pkgComponent.IComponent
	}
)

func (p *Page) Route() string {
	return p.route
}

func (p *Page) Layout() *pkgLayout.Layout {
	return p.layout
}

func (p *Page) Content() template.HTML {
	ctx := p.layout.Context()

	if p.inline_content != "" {
		ctx["content"] = p.inline_content
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

func NewPageFromString(route string, layout *pkgLayout.Layout, content string) *Page {
	return &Page{
		route:          route,
		layout:         layout,
		inline_content: content,
	}
}

func NewPageFromComponents(route string, layout *pkgLayout.Layout, components []pkgComponent.IComponent) *Page {
	return &Page{
		route:      route,
		layout:     layout,
		components: components,
	}
}
