package handlers

import (
	"net/http"

	"github.com/Brassalsa/m-chat/internal/db/schema"
	"github.com/Brassalsa/m-chat/internal/types"
	"github.com/Brassalsa/m-chat/pkg"
	"github.com/Brassalsa/m-chat/pkg/helpers"
	"github.com/Brassalsa/m-chat/pkg/res"
)

type UserHandler struct {
	types.Handler
}

type LoginUser struct {
	Email string `bson:"email"`
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var resCode = http.StatusOK

	formData := res.NewFormData()
	email := r.FormValue("email")
	password := r.FormValue("password")
	formData.Values["email"] = email
	formData.Values["password"] = password

	defer func() {
		if resCode == http.StatusFound {
			u.Redirect(w, r, "/home")
			return
		}
		pkg.RespondTempl(w, resCode, "log-in", formData)
	}()

	user := schema.User{}
	err := u.Db.Get(u.Coll, types.Email{
		Email: email,
	}, &user)

	if err != nil {
		formData.Errors["form"] = err.Error()
		resCode = http.StatusUnprocessableEntity
		return
	}

	if user.Email == "" {
		formData.Errors["email"] = "email not found"
		resCode = http.StatusUnprocessableEntity
		return
	}

	// check password
	if !pkg.CompareHash(user.Password, password) {
		formData.Errors["password"] = "wrong password"
		resCode = http.StatusUnprocessableEntity
		return
	}

	if err := helpers.GenerateAndSetTokens(w,
		types.Id{
			Id: user.Id,
		}); err != nil {
		formData.Errors["form"] = "error generating token"
		resCode = http.StatusUnprocessableEntity
		return
	}

	resCode = http.StatusFound
}

type RegisterUser struct {
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Password string `bson:"password"`
}

func (u *UserHandler) Regsiter(w http.ResponseWriter, r *http.Request) {
	var resCode = http.StatusFound

	formData := res.NewFormData()

	username := r.FormValue("username")
	formData.Values["username"] = username

	email := r.FormValue("email")
	formData.Values["email"] = email

	name := r.FormValue("name")
	formData.Values["name"] = name

	password := r.FormValue("password")
	formData.Values["password"] = password

	confirmPassword := r.FormValue("confirm_password")
	formData.Values["confirm_password"] = confirmPassword

	defer func() {
		if resCode == http.StatusFound {
			u.Redirect(w, r, "/home")
			return
		}
		pkg.RespondTempl(w, resCode, "register", formData)
	}()

	if helpers.CheckEmpty([]string{username, email, password}) {
		formData.Errors["form"] = "username, email and password are required"
		resCode = http.StatusUnprocessableEntity
		return
	}

	if password != confirmPassword {
		formData.Errors["confirm_password"] = "passwords don't match"
		resCode = http.StatusUnprocessableEntity
		return
	}

	user := schema.User{}

	if err := u.Db.Get(u.Coll,
		types.Username{Username: username},
		&user,
	); err != nil {
		formData.Errors["form"] = err.Error()
		resCode = http.StatusUnprocessableEntity
		return
	}

	if user.Username != "" && user.Username == username {
		formData.Errors["username"] = "username already taken"
		resCode = http.StatusUnprocessableEntity
		return
	}

	if user.Email != "" && user.Email == email {
		formData.Errors["email"] = "email already exists"
		resCode = http.StatusUnprocessableEntity
		return
	}

	if err := pkg.HashString(&password); err != nil {
		formData.Errors["form"] = "error hashing password"
		resCode = http.StatusUnprocessableEntity
		return
	}

	if err := u.Db.Add(u.Coll, RegisterUser{
		Username: username,
		Email:    email,
		Name:     name,
		Password: password,
	}); err != nil {
		formData.Errors["form"] = err.Error()
		resCode = http.StatusUnprocessableEntity
		return
	}

	// find inserted user
	if err := u.Db.Get(u.Coll, types.Email{
		Email: email,
	}, &user); err != nil {
		formData.Errors["form"] = err.Error()
		resCode = http.StatusUnprocessableEntity
		return
	}

	if err := helpers.GenerateAndSetTokens(w,
		types.Id{
			Id: user.Id,
		}); err != nil {
		formData.Errors["form"] = "error generating token"
		resCode = http.StatusUnprocessableEntity
		return
	}

	resCode = http.StatusFound
}
