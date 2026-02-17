package main

import (
	"net/http"
	"path"
)

func (app *App) render(w http.ResponseWriter, filename string, data any) {
	filePath := path.Join("./template/", filename)
	app.TemplateRender.Render(w, filePath, data)
}
