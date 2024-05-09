package main

import (
	"github.com/JensonCode/tentrek/api"
	"github.com/JensonCode/tentrek/internal/auth"
	"github.com/JensonCode/tentrek/internal/database"
	"github.com/JensonCode/tentrek/internal/server"
	"github.com/JensonCode/tentrek/internal/store"
)

func main() {

	db, err := database.InitPostgreSQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	auth.NewOAuth()

	server := server.NewServer()

	store := store.NewStore(db.DB)

	api.RegisterRoutes(server.Router, store)

	server.ListenAndServe()
}
