package configs

import (
	"fmt"
	"os"
)

type PostgresConf struct {
	Host     string
	User     string
	Password string
	DBName   string
}

func NewPostgresConf() string {
	conf := &PostgresConf{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
	}

	if conf.Host == "host.docker.internal" {
		return fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.User, conf.Password, conf.DBName)
	}

	return fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s", conf.Host, conf.User, conf.Password, conf.DBName)
}
