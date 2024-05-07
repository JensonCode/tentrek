package configs

import "os"

type ServerConf struct {
	Port string
}

func NewServerConf() *ServerConf {
	return &ServerConf{
		Port: os.Getenv("API_SERVER_PORT"),
	}
}
