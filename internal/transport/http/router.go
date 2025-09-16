package http

import (
	"app/internal/transport/http/category"
	"app/internal/transport/http/product"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(CategoryHandler *category.Handler, ProductHandler *product.Handler) http.Handler {
	r := chi.NewRouter()

	// маршруты для категории
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", CategoryHandler.GetCategories)
		r.Get("/{id}", CategoryHandler.GetCategory)
		// r.Post("/", userHandler.Create)
	})

	r.Route("/products", func(r chi.Router) {
		r.Get("/", ProductHandler.GetList)
		// r.Get("/{id}", CategoryHandler.GetCategory)
		r.Post("/", ProductHandler.Create)
		r.Put("/{id}", ProductHandler.Update)
		r.Patch("/{id}", ProductHandler.Update)
	})

	// // маршруты для заказов
	// r.Route("/orders", func(r chi.Router) {
	//     r.Get("/{id}", orderHandler.GetOrder) // GET /orders/1
	//     r.Post("/", orderHandler.Create)      // POST /orders
	// })

	r.Mount("/api", r)

	return r
}
