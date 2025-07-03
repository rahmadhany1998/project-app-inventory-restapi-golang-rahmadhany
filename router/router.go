package router

import (
	"project-app-inventory-restapi-golang-rahmadhany/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", h.UserHandler.GetAll)
		r.Post("/", h.UserHandler.Create)
		r.Get("/{id}", h.UserHandler.GetByID)
		r.Put("/{id}", h.UserHandler.Update)
		r.Delete("/{id}", h.UserHandler.Delete)
	})

	return r
}
