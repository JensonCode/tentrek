package api

import (
	"net/http"

	"github.com/JensonCode/tentrek/internal/store"
	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func Serve(h HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := MiddlewareChain(h, middlewares...)(w, r); err != nil {
			WriteError(w, err)
		}
	}
}

type Handler struct {
	store *store.Store
}

func RegisterRoutes(router *mux.Router, store *store.Store) {
	h := &Handler{store: store}

	router.HandleFunc("/auth/user/{service}", Serve(h.AuthHandlers))

	//OAuth
	router.HandleFunc("/auth/{provider}", Serve(h.OAuthLogin))
	router.HandleFunc("/auth/{provider}/callback", Serve(h.OAuthCallback))
	router.HandleFunc("/logout/{provider}", Serve(h.OAuthLogout))
}
