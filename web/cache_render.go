package main

import (
	"html/template"
	"net/http"
	"sync"
)

type TemplateRender struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	isDev       bool
	templateDir string
}

func NewTemplateRender(isDev bool, templateDir string) *TemplateRender {
	return &TemplateRender{
		cache:       make(map[string]*template.Template),
		isDev:       isDev,
		templateDir: templateDir,
	}
}

func (tr *TemplateRender) Render(w http.ResponseWriter, filename string, data any) {
	tr.mutex.RLock()
	temp, ok := tr.cache[filename]
	tr.mutex.RUnlock()

	if ok && !tr.isDev {
		err := temp.Execute(w, data)
		if err != nil {
			panic(err)
		}
		return
	}

	temp, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}

	if !tr.isDev {
		tr.mutex.Lock()
		tr.cache[filename] = temp
		tr.mutex.Unlock()
	}

	err = temp.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
