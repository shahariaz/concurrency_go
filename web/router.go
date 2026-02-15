package main

import "net/http"

func (app *App) SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/create-user", app.UserHandler.CreateUser)
	return mux
}
