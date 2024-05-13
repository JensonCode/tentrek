package middleware

import (
	"net/http"
	"os"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(h http.HandlerFunc, mws ...Middleware) http.HandlerFunc {

	for _, mw := range mws {
		h = mw(h)
	}

	return h
}

func WithCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		allowedOrigin := os.Getenv("APP_BASE_URL")

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}
