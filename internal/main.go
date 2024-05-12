package main

import (
	"github.com/JensonCode/tentrek/api"
	"github.com/JensonCode/tentrek/oauth"

	"github.com/JensonCode/tentrek/database"
	"github.com/JensonCode/tentrek/store"
)

func main() {

	db, err := database.InitPostgreSQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	store := store.NewStore(db.DB)

	oauth.NewOAuth()

	api.Run(store)

}
