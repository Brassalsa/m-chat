package pkg

import (
	"html/template"
	"io"
	"strings"
)

type Templates struct {
	templates *template.Template
}

type LayoutPage struct {
	Content string
}

func NewTemplate() *Templates {
	templ, err := template.ParseGlob("*web/views/*.html")
	if err != nil {
		if strings.Contains(err.Error(), "pattern matches no files") {
			templ, err = template.ParseGlob("../web/views/*.html")
			if err != nil {
				panic(err)
			}

		} else {
			panic(err)
		}
	}
	return &Templates{
		templates: templ,
	}
}

func (t *Templates) Render(wr io.Writer, name string, payload interface{}) error {
	return t.templates.ExecuteTemplate(wr, name, payload)
}

func (t *Templates) RenderWLayout(wr io.Writer, layout, name string, payload interface{}) error {

	if err := t.Render(wr, layout, payload); err != nil {
		return err
	}

	return t.Render(wr, name, payload)
}
