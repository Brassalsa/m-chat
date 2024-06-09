package api

import (
	"net/http"

	"github.com/Brassalsa/m-chat/api/handlers"
	"github.com/Brassalsa/m-chat/pkg"
)

type Data struct {
	Payload string
}

func HandleV1Route() http.Handler {
	r := pkg.NewRouter()
	layout := "layout"
	r.RegisterHandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Payload: "Hello from server",
		}
		pkg.RespondPage(w, 200, layout, "index", data)
	})
	r.RegisterHandleFunc("GET /sign-in", func(w http.ResponseWriter, r *http.Request) {
		pkg.RespondPage(w, 200, layout, "sign-in", "")
	})

	userHandler := handlers.UserHandler{
		Rmux: r.RMux,
	}
	r.RegisterHandleFunc("POST /sign-in", userHandler.Login)
	return r.RMux
}
