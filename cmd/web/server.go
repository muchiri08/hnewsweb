package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *application) listenAndServe() error {
	host := fmt.Sprintf("%s:%s", a.server.host, a.server.port)

	server := http.Server{
		Handler:      a.routes(),
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	a.infoLog.Printf("Server listening on %s\n", host)

	return server.ListenAndServe()
}
