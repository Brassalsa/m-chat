package middlewares

import (
	"context"
	"net/http"

	"github.com/Brassalsa/m-chat/internal/db"
	"github.com/Brassalsa/m-chat/internal/db/schema"
	"github.com/Brassalsa/m-chat/internal/types"
	"github.com/Brassalsa/m-chat/pkg"
	"github.com/Brassalsa/m-chat/pkg/helpers"
	"github.com/Brassalsa/m-chat/pkg/res"
)

type AuthHandler struct {
	types.Handler
}

func (h *AuthHandler) OnlyAuth(fn types.HandleFunc) types.HandleFunc {

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
		user := schema.User{}
		if err := h.Db.Get("users", types.Id{
			Id: id.Id,
		}, &user); err != nil {
			if err.Error() == db.NOT_FOUND {
				pkg.RespondPage(w, 200, "layout", "error", res.ErrorData{
					StatusCode: 404,
					Message:    "User not found",
				})
				return
			}
			pkg.RespondPage(w, 200, "layout", "error", res.ErrorData{
				StatusCode: 401,
				Message:    "Unauthorized please login to continue",
			})
			return
		}

		newR := r.WithContext(context.WithValue(h.Ctx, types.Id{}, id))
		fn(w, newR)
	}
}
