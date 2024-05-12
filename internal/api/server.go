package api

import (
	"net/http"

	"github.com/JensonCode/tentrek/configs"

	"github.com/JensonCode/tentrek/store"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	port   string
	store  *store.Store
}

func Run(store *store.Store) {
	conf := configs.NewServerConf()

	mux := mux.NewRouter()

	server := &Server{
		router: mux,
		port:   conf.Port,
		store:  store,
	}

	server.RegisterRoutes()

	http.ListenAndServe(server.port, server.router)
}
