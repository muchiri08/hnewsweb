package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()

	//Middlewares stack from chi
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	//my middleware
	mux.Use(a.loadSession)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	//create file server
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return mux
}
