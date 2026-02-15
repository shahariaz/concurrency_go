package main

import (
	"html/template"
	"net/http"
	"path"
)

func (app *App) render(w http.ResponseWriter, filename string, data any) {
	filePath := path.Join("./template/", filename)

	temp, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
