package api

import (
	"net/http"
	"os"
)

type Middleware func(Handler) Handler

func MiddlewareChain(h Handler, middlewares ...Middleware) Handler {
	for _, m := range middlewares {
		h = m(h)
	}

	return h
}

func WithCors(h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {

		allowedOrigin := os.Getenv("APP_BASE_URL")

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return nil
		}

		return h(w, r)
	}
}
