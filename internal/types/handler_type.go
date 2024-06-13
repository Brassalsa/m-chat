package types

import (
	"net/http"

	"github.com/Brassalsa/m-chat/internal/db"
)

type Handler struct {
	Rmux *http.ServeMux
	Db   *db.MongoDb
	Coll string
}
