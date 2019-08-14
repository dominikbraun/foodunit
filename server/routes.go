// Package server provides functions to control FoodUnit's HTTP Server.
package server

import (
	"github.com/dominikbraun/foodunit/controllers/http"
	"github.com/go-chi/chi"
)

// mountRoutes builds all API routes and registers the responsible handlers.
func (srv *Server) mountRoutes() {
	r := chi.NewRouter()

	r.Route("/offers", func(r chi.Router) {
		r.Post("/new", http.PostOffer)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.GetOffer)
			r.Put("/", http.PutOffer)
			r.Delete("/", http.DeleteOffer)

			r.Route("/orders", func(r chi.Router) {
				r.Get("/", http.GetOrders)
			})

			r.Route("/order", func(r chi.Router) {
				r.Post("/", http.F)
				r.Get("/", http.F)
				r.Put("/", http.F)
			})
		})
	})

	r.Route("/suppliers", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.F)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", http.F)
		r.Post("/login", http.F)
		r.Get("/logout", http.F)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.F)
		})
	})

	srv.router.Mount("/", r)
}
