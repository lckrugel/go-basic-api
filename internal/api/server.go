package api

import (
	"log"
	"net/http"

	"github.com/lckrugel/go-stock/internal/config"
)

type ServerHTTP struct {
	server *http.Server
}

func NewHTTPServer(cfg config.AppConfig) *ServerHTTP {
	router := http.NewServeMux()
	routes(router)
	return &ServerHTTP{
		server: &http.Server{
			Addr:    cfg.Host + ":" + cfg.Port,
			Handler: router,
		},
	}
}

func (http *ServerHTTP) Start() error {
	log.Println("Server starting on", http.server.Addr)
	return http.server.ListenAndServe()
}
