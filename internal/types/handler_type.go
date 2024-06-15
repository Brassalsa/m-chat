package types

import (
	"net/http"

	"github.com/Brassalsa/m-chat/pkg/persister"
)

type Handler struct {
	Rmux *http.ServeMux
	Db   persister.Persister
	Coll string
}
