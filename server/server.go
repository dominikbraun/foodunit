// Package server provides functions to control FoodUnit's HTTP Server.
package server

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Start creates and starts new server instance which will expose
// the REST API. Ctrl + C will attempt a graceful shutdown.
func Start() {
	r := newRouter()
	r.Mount("/", routeTree())

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		<-interrupt
		ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)

		err := srv.Shutdown(ctx)
		if err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
		defer cancelFn()
	}()

	log.Fatal(srv.ListenAndServe())
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)
	return r
}
