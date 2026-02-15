package main

import "net/http"

func (app *App) serve() error {
	server := http.Server{
		Addr:    ":8080",
		Handler: app.SetupRouter(),
	}
	err := server.ListenAndServe()
	return err
}
