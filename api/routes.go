package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(http.ResponseWriter, *http.Request) error

func Serve(h Handler, middlewares ...Middleware) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := MiddlewareChain(h, middlewares...)(w, r); err != nil {
			WriteError(w, err)
		}
	}
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", Serve(handleHello, WithCors))
}

func handleHello(w http.ResponseWriter, r *http.Request) error {
	return WriteResponse(w, 200, "hello")
}
