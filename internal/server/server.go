package server

import (
	"log"
	"net/http"

	"github.com/JensonCode/tentrek/internal/configs"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Port   string
}

func NewServer() *Server {
	conf := configs.NewServerConf()

	mux := mux.NewRouter()

	return &Server{
		Router: mux,
		Port:   conf.Port,
	}
}

func (s *Server) ListenAndServe() error {
	log.Println("TenTrek api server listening")

	return http.ListenAndServe(s.Port, s.Router)
}
