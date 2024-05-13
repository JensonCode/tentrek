package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/JensonCode/tentrek/jwt"
)

func writeUnauthorized(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func WithAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 || strings.ToLower(tokenString[:6]) != "bearer" {
			writeUnauthorized(w, errors.New("authorization header incorrect"))
			return
		}

		_, err := jwt.ValidateToken(tokenString[7:])
		if err != nil {
			writeUnauthorized(w, errors.New("invalid jwt token"))
			return
		}

		h(w, r)
	}
}
