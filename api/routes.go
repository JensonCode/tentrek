package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(http.ResponseWriter, *http.Request) error

func Serve(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			WriteError(w, err)
		}
	}
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", Serve(func(w http.ResponseWriter, r *http.Request) error {
		return WriteResponse(w, 200, "hello")
	}))
}
