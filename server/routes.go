// Package server provides functions to control FoodUnit's HTTP Server.
package server

import (
	"github.com/dominikbraun/foodunit/controllers/http"
	"github.com/go-chi/chi"
)

// routeTree builds all API routes and registers the responsible handlers.
func routeTree() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/offers", func(r chi.Router) {
		r.Post("/new", http.PostOffer)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.GetOffer)
			r.Put("/", http.PutOffer)
			r.Delete("/", http.DeleteOffer)

			r.Route("/orders", func(r chi.Router) {
				r.Get("/", http.GetOrders)
				r.Put("/", http.PutOrder)

			})
		})
	})

	return r
}
