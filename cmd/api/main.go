package main

import (
	"github.com/JensonCode/tentrek/api"
	"github.com/JensonCode/tentrek/internal/database"
	"github.com/JensonCode/tentrek/internal/server"
)

func main() {

	if _, err := database.InitPostgreSQL(); err != nil {
		panic(err)
	}

	server := server.NewServer()

	api.RegisterRoutes(server.Router)

	server.ListenAndServe()
}
