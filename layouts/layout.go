package layouts

import (
	"fmt"
	"html/template"
	"log"
	"time"
)

type ILayout interface {
	Template() *template.Template
	Context() map[string]string
}

type Layout struct {
	template *template.Template
	context  map[string]string
}

func NewTemplateFromFilepath(filepath string) *template.Template {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func NewTemplateFromString(inline string) *template.Template {
	name := fmt.Sprintf("%d", time.Now().Unix())
	t, err := template.New(name).Parse(inline)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func NewLayout(t *template.Template, ctx map[string]string) *Layout {
	return &Layout{
		template: t,
		context:  ctx,
	}
}

func (l *Layout) Template() *template.Template {
	return l.template
}

func (l *Layout) Context() map[string]string {
	return l.context
}
