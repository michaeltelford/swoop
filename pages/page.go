package pages

import (
	"bytes"
	"fmt"

	pkgComponents "github.com/michaeltelford/swoop/components"
	pkgLayouts "github.com/michaeltelford/swoop/layouts"
)

type IPage interface {
	Route() string              // The page URL path e.g. "/foo/bar"
	Layout() *pkgLayouts.Layout // The layout to use for the page
	Content() string            // The content of the page
}

func NewPageFromString(route string, layout *pkgLayouts.Layout, content string) *Page {
	return &Page{
		route:   route,
		layout:  layout,
		content: content,
	}
}

func NewPageFromComponents(name, route string, layout *pkgLayouts.Layout, components []pkgComponents.IComponent) *Page {
	return &Page{
		route:      route,
		layout:     layout,
		components: components,
	}
}

type Page struct {
	route      string
	layout     *pkgLayouts.Layout
	content    string
	components []pkgComponents.IComponent
}

func (p *Page) Route() string {
	return p.route
}

func (p *Page) Layout() *pkgLayouts.Layout {
	return p.layout
}

func (p *Page) Content() string {
	p.renderLayout()

	if p.content != "" {
		//
	} else if p.components != nil {
		// TODO.
		componentContent := p.renderComponents()
		fmt.Printf("Component content: %s\n", componentContent)
	}

	return p.content
}

func (p *Page) renderLayout() {
	tmpl := p.layout.Template()
	var buffer bytes.Buffer
	tmpl.ExecuteTemplate(&buffer, tmpl.Name(), p.layout.Context())
	p.content = buffer.String()
}

func (p *Page) renderComponents() string {
	return "TODO"
}
