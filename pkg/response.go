package pkg

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

var templ = NewTemplate()

func RespondTempl(w http.ResponseWriter, code int, name string, payload interface{}) {
	w.WriteHeader(code)

	if err := templ.Render(w, name, payload); err != nil {
		log.Printf("error rendering %s template, %s\n", name, err)
	}
}

func RespondPage(w http.ResponseWriter, code int, layout, name string, payload interface{}) {
	w.WriteHeader(code)

	if err := templ.RenderWLayout(w, layout, name, payload); err != nil {
		log.Printf("error rendering %s template, %s\n", name, err)
	}
}

func RespondBytes(w io.Writer, name string, payload interface{}) {
	buf := new(bytes.Buffer)

	if err := templ.Render(buf, name, payload); err != nil {
		log.Printf("error rendering %s template, %s\n", name, err)
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("error sending data")
	}
}
