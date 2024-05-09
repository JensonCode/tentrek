package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JensonCode/tentrek/internal/configs"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	DB *sql.DB
}

func InitPostgreSQL() (*PostgreSQL, error) {
	log.Println("** Connecting to Postgres DB **")

	connectString := configs.NewPostgresConf()

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, fmt.Errorf("open postgres failed:%s", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping postgres failed:%s", err)
	}

	psql := &PostgreSQL{DB: db}

	return psql, nil
}
