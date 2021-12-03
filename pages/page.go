package pages

type IPage interface {
	Route() string   // e.g. /foo/bar
	Content() string // e.g. The content of the page
}

func NewPage(route, content string) *Page {
	return &Page{
		route:   route,
		content: content,
	}
}

type Page struct {
	route   string
	content string
}

func (p *Page) Route() string {
	return p.route
}

func (p *Page) Content() string {
	return p.content
}
