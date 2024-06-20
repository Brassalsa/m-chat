package types

import (
	"context"
	"net/http"

	"github.com/Brassalsa/m-chat/pkg/persister"
)

// handler type
type Handler struct {
	Rmux *http.ServeMux
	Db   persister.Persister
	Coll string
	Ctx  context.Context
}

func (h *Handler) RedirectWCode(w http.ResponseWriter, r *http.Request, to string, code int) {
	w.Header().Set("HX-Redirect", to)
	w.WriteHeader(code)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request, to string) {
	h.RedirectWCode(w, r, to, http.StatusFound)
}

type HandleFunc func(w http.ResponseWriter, r *http.Request)
