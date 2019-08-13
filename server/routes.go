// Package server provides functions to control FoodUnit's HTTP Server.
package server

import (
	"github.com/dominikbraun/foodunit/controllers/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// routes builds the API routes and returns the app's router.
func routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)

	r.Route("/offers", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.GetOffer)
			r.Put("/", http.PutOffer)

			r.Route("/orders", func(r chi.Router) {
				r.Get("/", http.GetOrders)
				r.Put("/", http.PutOrder)
			})
		})
	})

	return r
}
