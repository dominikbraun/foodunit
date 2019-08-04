package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func routeHandlers() *chi.Mux {
	r := chi.NewRouter()
	
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Route("/v1", func (r chi.Router) {
		// r.Mount("/user", user.Routes())
		// r.Mount("/offer", offer.Routes())
		// r.Mount("/supplier", supplier.Routes())
		// r.Mount("/dish", dish.Routes())
		// r.Mount("/order", order.Routes())
	})

	return r
}

func Start() error {
	r := chi.NewRouter()
	
	walk := func(method string, route string, handle http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walk); err != nil {
		return err
	}
	
	if err := http.ListenAndServe(":8000", r); err != nil {
		return err
	}
	return nil
}
