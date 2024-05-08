package main

import (
	"net/http"

	"github.com/JensonCode/tentrek/internal/database"
	"github.com/JensonCode/tentrek/internal/server"
)

func main() {

	if _, err := database.InitPostgreSQL(); err != nil {
		panic(err)
	}

	server := server.NewServer()

	server.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server.ListenAndServe()
}
