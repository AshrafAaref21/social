package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ashrafaaref20/social/internal/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route(
		"/v1", func(r chi.Router){
			r.Get("/health", app.healthCheckHandler)
		},
	)
	r.Get("/health", app.healthCheckHandler)

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
