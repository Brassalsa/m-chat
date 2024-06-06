package main

import (
	"log"
	"net/http"
)

type Page struct {
	Payload string
}

func main() {
	s := newServer(":3000", "static")

	s.RegisterHandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Page{
			Payload: "hello world",
		}
		respondTempl(w, 200, "index", data)
	})

	log.Printf("serving @ http://loclahost:%s\n", s.ListenAddr)
	log.Fatal(s.Listen())
}
