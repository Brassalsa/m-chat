package main

import "net/http"

type Router struct {
	RMux *http.ServeMux
}

func newRouter() *Router {
	return &Router{
		RMux: http.NewServeMux(),
	}
}

func (r *Router) RegisterHandleFunc(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) {
	r.RMux.HandleFunc(pattern, handleFunc)
}

func (r *Router) RegisterHandler(pattern string, handler http.Handler) {
	r.RMux.Handle(pattern, handler)
}
