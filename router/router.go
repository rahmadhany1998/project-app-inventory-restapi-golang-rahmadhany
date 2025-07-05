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

	r.Route("/api/v1/categories", func(r chi.Router) {
		r.Get("/", h.CategoryHandler.GetAll)
		r.Post("/", h.CategoryHandler.Create)
		r.Get("/{id}", h.CategoryHandler.GetByID)
		r.Put("/{id}", h.CategoryHandler.Update)
		r.Delete("/{id}", h.CategoryHandler.Delete)
	})

	r.Route("/api/v1/racks", func(r chi.Router) {
		r.Get("/", h.RackHandler.GetAll)
		r.Post("/", h.RackHandler.Create)
		r.Get("/{id}", h.RackHandler.GetByID)
		r.Put("/{id}", h.RackHandler.Update)
		r.Delete("/{id}", h.RackHandler.Delete)
	})

	r.Route("/api/v1/warehouses", func(r chi.Router) {
		r.Get("/", h.WarehouseHandler.GetAll)
		r.Post("/", h.WarehouseHandler.Create)
		r.Get("/{id}", h.WarehouseHandler.GetByID)
		r.Put("/{id}", h.WarehouseHandler.Update)
		r.Delete("/{id}", h.WarehouseHandler.Delete)
	})

	r.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", h.ProductHandler.GetAll)
		r.Post("/", h.ProductHandler.Create)
		r.Get("/{id}", h.ProductHandler.GetByID)
		r.Put("/{id}", h.ProductHandler.Update)
		r.Delete("/{id}", h.ProductHandler.Delete)
	})

	r.Route("/api/v1/sales", func(r chi.Router) {
		r.Get("/", h.SaleHandler.GetAll)
		r.Post("/", h.SaleHandler.Create)
		r.Get("/{id}", h.SaleHandler.GetByID)
	})

	r.Route("/api/v1/report", func(r chi.Router) {
		r.Get("/summary", h.SaleHandler.GetReportSummaryByDate)
	})

	return r
}
