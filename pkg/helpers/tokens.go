package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Brassalsa/m-chat/pkg"
)

const defaultTokenName = "Bearer"

func GenerateAndSetTokens(w http.ResponseWriter, payload interface{}) error {
	token, err := pkg.GenerateJWT(payload)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     defaultTokenName,
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), // Set an appropriate expiration time
		HttpOnly: true,                           // Ensure the cookie is not accessible via JavaScript
	})

	return nil
}

func DecodeToken(r *http.Request, decodeTo interface{}) error {
	// get toke from header
	token := r.Header.Get(defaultTokenName)

	if token == "" {
		if ck, err := r.Cookie(defaultTokenName); err == nil {
			token = ck.Value
		}
	}

	if token == "" {
		return fmt.Errorf(`"%s" not found `, defaultTokenName)
	}
	if err := pkg.ValidateToken(token, decodeTo); err != nil {
		return err
	}
	return nil
}
