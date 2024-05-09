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

	if err := psql.migrate(); err != nil {
		return nil, err
	}

	return psql, nil
}

func (p *PostgreSQL) Close() {
	log.Println("** Disconnecting Postgres DB **")
	p.DB.Close()
}

func (p *PostgreSQL) migrate() error {
	log.Println("** Migrate tables **")
	for _, query := range migration {
		_, err := p.DB.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to migrate: %s\n query: %q", err, query)
		}
	}
	return nil
}
