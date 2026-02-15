package main

import (
	"html/template"
	"net/http"
)

var htmlContent = `

<!DOCTYPE html>
<html lang="en">
<head>
			
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Go Web Server</title>
	</head>
	<body>
		<h1>Welcome to the Go Web Server!</h1>
		<p>This is the {{.Page}} page.</p>
		</body>
		</html>
		`

type PageData struct {
	Page string
}

func RenderTemplate(w http.ResponseWriter, data PageData) {
	template, err := template.New("web-page").Parse(htmlContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "index.html", PageData{Page: "Home"})
}

func (app *App) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, "index.html", PageData{Page: "About"})

}

func contact(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{Page: "Contact"})
}
