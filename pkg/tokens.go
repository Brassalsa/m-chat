package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var defSec = "wBg5867Y+QgPDUV9ICOVv9A5PipeTiNQFMtIH2b+1jo="

func getSecret() []byte {
	var secretStr = os.Getenv("JWT_SECRET")
	if secretStr == "" {
		log.Println("JWT_SECRET is not in env, using defaults")
		return []byte(defSec)
	}
	return []byte(secretStr)
}

func GenerateJWT(payload interface{}) (string, error) {
	payloadJson, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("err in marshaling: %s", err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = string(payloadJson)
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24).Unix()

	tokenString, err := token.SignedString(getSecret())

	if err != nil {
		err = fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string, decodeTo interface{}) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return getSecret(), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userJson := claims["user"].(string)

		err := json.Unmarshal([]byte(userJson), &decodeTo)
		if err != nil {
			return fmt.Errorf("error in unmarshaling: %v", err)
		}

		return nil
	} else {
		return err
	}
}
