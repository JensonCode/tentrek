package configs

import (
	"fmt"
	"os"
)

type PostgreSQLConf struct {
	Host     string
	User     string
	Password string
	DBName   string
}

func NewPostgresConnString() string {
	conf := &PostgreSQLConf{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
	}

	return fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.User, conf.Password, conf.DBName)
}
