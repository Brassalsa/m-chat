package main

import (
	"bytes"
	"log"
	"net/http"
)

var templ = newTemplate()

func respondTempl(w http.ResponseWriter, code int, name string, payload interface{}) {
	w.WriteHeader(code)

	if err := templ.Render(w, name, payload); err != nil {
		log.Printf("error rendering %s template, %s\n", name, err)
	}
}

func respondBytes(w http.ResponseWriter, code int, name string, payload interface{}) {
	buf := new(bytes.Buffer)
	w.WriteHeader(code)

	if err := templ.Render(buf, name, payload); err != nil {
		log.Printf("error rendering %s template, %s\n", name, err)
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("error sending data")
	}
}
