package main

import (
	"net/http"
)

type Data struct {
	Payload string
}

func handleV1Route() http.Handler {
	r := newRouter()
	r.RegisterHandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Payload: "Hello from server",
		}
		respondTempl(w, 200, "index", data)
	})
	r.RegisterHandleFunc("GET /sign-in", func(w http.ResponseWriter, r *http.Request) {
		respondTempl(w, 200, "sign-in", "")
	})
	return r.RMux
}
