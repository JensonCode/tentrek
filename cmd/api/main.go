package main

import (
	"github.com/JensonCode/tentrek/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}

	server := server.NewServer()
	server.ListenAndServe()
}
