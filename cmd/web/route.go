package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(a.appName))
	})

	mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Comments"))
	})

	return mux
}
