package pkg

import (
	"net/http"
	"strings"
)

type Server struct {
	ListenAddr string
	StaticDir  string
	RMux       *http.ServeMux
}

func NewServer(listenAddr, staticDir string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		StaticDir:  staticDir,
		RMux:       http.NewServeMux(),
	}
}

func (s *Server) Listen() error {
	if s.StaticDir != "" {
		// register static file handler
		dir := formatDir(s.StaticDir)

		fs := http.FileServer(http.Dir(s.StaticDir))
		s.RegisterHandler(dir, http.StripPrefix(dir, fs))
	}

	return http.ListenAndServe(s.ListenAddr, s.RMux)
}

func (s *Server) RegisterHandleFunc(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) {
	s.RMux.HandleFunc(pattern, handleFunc)
}

func (s *Server) RegisterHandler(pattern string, handler http.Handler) {
	s.RMux.Handle(pattern, handler)
}

func formatDir(dir string) string {
	str := []string{}
	if dir[0] != '/' {
		str = append(str, "/")
	}
	str = append(str, dir)
	if dir[len(dir)-1] != '/' {
		str = append(str, "/")
	}

	return strings.Join(str, "")
}
