package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()

	//Middlewares stack from chi
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Comments"))
	})

	return mux
}
