package api

import (
	"net/http"

	"github.com/lckrugel/go-basic-api/internal/api/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	registerRoutes(mux)

	router := registerGlobalMiddlewares(mux)

	return router
}

func registerGlobalMiddlewares(router http.Handler) http.Handler {
	router = middleware.Logging(router)

	return router
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "hello": "world" }`))
	})
}
