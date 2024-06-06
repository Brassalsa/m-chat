package main

import (
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("*views/*.html")),
	}
}

func (t *Templates) Render(wr io.Writer, name string, payload interface{}) error {
	return t.templates.ExecuteTemplate(wr, name, payload)
}
