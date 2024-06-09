package handlers

import (
	"fmt"
	"net/http"

	"github.com/Brassalsa/m-chat/pkg"
)

type UserHandler struct {
	Rmux *http.ServeMux
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Printf("email %s, password %s\n", email, password)
	pkg.RespondTempl(w, 200, "sign-in", "")

}
