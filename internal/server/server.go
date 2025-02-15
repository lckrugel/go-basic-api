package server

import (
	"context"
	"log"
	"net/http"

	"github.com/lckrugel/go-basic-api/internal/config"
)

type ServerHTTP struct {
	server *http.Server
}

func NewHTTPServer(cfg config.AppConfig, h http.Handler) *ServerHTTP {
	return &ServerHTTP{
		server: &http.Server{
			Addr:    cfg.Host + ":" + cfg.Port,
			Handler: h,
		},
	}
}

func (s *ServerHTTP) Start() error {
	log.Println("Server starting on", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *ServerHTTP) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
