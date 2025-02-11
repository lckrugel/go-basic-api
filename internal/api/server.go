package api

import (
	"context"
	"log"
	"net/http"

	"github.com/lckrugel/go-basic-api/internal/config"
)

type ServerHTTP struct {
	server *http.Server
}

func NewHTTPServer(cfg config.AppConfig) *ServerHTTP {
	router := NewRouter()

	return &ServerHTTP{
		server: &http.Server{
			Addr:    cfg.Host + ":" + cfg.Port,
			Handler: router,
		},
	}
}

func (s *ServerHTTP) Start() error {
	log.Println("Server starting on", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *ServerHTTP) Close(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
