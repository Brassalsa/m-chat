package middlewares

import (
	"context"
	"net/http"

	"github.com/Brassalsa/m-chat/internal/db"
	"github.com/Brassalsa/m-chat/internal/types"
	"github.com/Brassalsa/m-chat/pkg"
	"github.com/Brassalsa/m-chat/pkg/helpers"
	"github.com/Brassalsa/m-chat/pkg/res"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func OnlyAuth(c context.Context, db *db.MongoDb, fn Handler) Handler {

	return func(w http.ResponseWriter, r *http.Request) {
		id := types.Id{}
		if err := helpers.DecodeToken(r, &id); err != nil {
			pkg.RespondPage(w, 200, "layout", "error", res.ErrorData{
				StatusCode: 401,
				Message:    "Unauthorized please login to continue",
			})
			return
		}

		// check user in db

		newR := r.WithContext(context.WithValue(c, types.Id{}, id))
		fn(w, newR)
	}
}
