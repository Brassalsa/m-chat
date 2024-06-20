package api

import (
	"context"
	"log"
	"net/http"

	"github.com/Brassalsa/m-chat/internal/api/handlers"
	"github.com/Brassalsa/m-chat/internal/api/middlewares"
	"github.com/Brassalsa/m-chat/internal/db"
	"github.com/Brassalsa/m-chat/internal/types"
	"github.com/Brassalsa/m-chat/pkg"
	"github.com/Brassalsa/m-chat/pkg/res"
)

type Data struct {
	Payload string
}

func HandleV1Route(ctx context.Context, dbC *db.MongoDb) http.Handler {
	r := pkg.NewRouter()
	layout := "layout"
	dbC.AddCollection([]string{"users"})

	r.RegisterHandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Payload: "Hello from server",
		}
		pkg.RespondPage(w, 200, layout, "index", data)
	})
	r.RegisterHandleFunc("GET /sign-in", func(w http.ResponseWriter, r *http.Request) {
		formData := res.NewFormData()
		values := r.URL.Query()
		registerValue := values.Get("register")
		if len(registerValue) == 0 || registerValue == "0" {
			formData.Values["templ"] = "log-in"
		} else {
			formData.Values["templ"] = "register"
		}

		pkg.RespondPage(w, 200, layout, "sign-in", formData)
	})

	userH := handlers.UserHandler{
		Handler: types.Handler{
			Rmux: r.RMux,
			Db:   dbC,
			Coll: "users",
			Ctx:  ctx,
		},
	}

	r.RegisterHandleFunc("GET /log-in", func(w http.ResponseWriter, r *http.Request) {
		pkg.RespondTempl(w, 200, "log-in", res.NewFormData())
	})
	r.RegisterHandleFunc("GET /register", func(w http.ResponseWriter, r *http.Request) {
		pkg.RespondTempl(w, 200, "register", res.NewFormData())
	})
	r.RegisterHandleFunc("POST /log-in", userH.Login)
	r.RegisterHandleFunc("POST /register", userH.Regsiter)

	authH := middlewares.AuthHandler{
		Handler: types.Handler{
			Rmux: r.RMux,
			Db:   dbC,
			Coll: "users",
			Ctx:  ctx,
		},
	}

	r.RegisterHandleFunc("GET /home", authH.OnlyAuth(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(types.Id{}).(types.Id)
		r.Context().Done()
		log.Println("id ", id)

		pkg.RespondPage(w, 200, layout, "home", "bruh")
	}))

	return r.RMux
}
