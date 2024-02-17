package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if !validateToken(authHeaderParts[1]) {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func validateToken(secretToken string) bool {
	mySigningKey := []byte("mysuppersecret")
	token, err := jwt.Parse(secretToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}
