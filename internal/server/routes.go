package server

import (
	"net/http"

	"github.com/lckrugel/go-basic-api/internal/app/controller"
)

type Controllers struct {
	UserController *controller.UserController
}

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	return &Router{
		mux,
	}
}

func (r *Router) RegisterRoutes(c *Controllers) {
	r.Mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "hello": "world" }`))
	})

	r.Mux.HandleFunc("GET /users", c.UserController.List)
	r.Mux.HandleFunc("POST /users", c.UserController.Create)
	r.Mux.HandleFunc("PUT /users/{id}", c.UserController.Update)
	r.Mux.HandleFunc("DELETE /users/{id}", c.UserController.Delete)
	r.Mux.HandleFunc("GET /users/{id}", c.UserController.Show)
}
