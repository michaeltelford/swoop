package page

import (
	"bytes"
	"html/template"
	"log"

	pkgComponents "github.com/michaeltelford/swoop/component"
	pkgLayouts "github.com/michaeltelford/swoop/layout"
)

type (
	IPage interface {
		Route() string              // The page URL path e.g. "/foo/bar"
		Layout() *pkgLayouts.Layout // The layout to use for the page
		Content() string            // The content of the page
	}

	Page struct {
		route          string
		layout         *pkgLayouts.Layout
		inline_content string
		components     []pkgComponents.IComponent
	}
)

func (p *Page) Route() string {
	return p.route
}

func (p *Page) Layout() *pkgLayouts.Layout {
	return p.layout
}

func (p *Page) Content() string {
	var content string

	if p.inline_content != "" {
		ctx := p.layout.Context()
		ctx["inline_content"] = p.inline_content
		content = parseTemplate(p.layout.Template(), ctx)
	} else if p.components != nil {
		log.Fatal("TODO: Page.Content() called on page with components but no content")
	} else {
		log.Fatal("No content or components to render")
	}

	return content
}

func parseTemplate(tmpl *template.Template, ctx map[string]string) string {
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, ctx); err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}

func NewPageFromString(route string, layout *pkgLayouts.Layout, content string) *Page {
	return &Page{
		route:          route,
		layout:         layout,
		inline_content: content,
	}
}

func NewPageFromComponents(route string, layout *pkgLayouts.Layout, components []pkgComponents.IComponent) *Page {
	return &Page{
		route:      route,
		layout:     layout,
		components: components,
	}
}
