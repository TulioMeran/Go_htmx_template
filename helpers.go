package main

import (
	"html/template"
	"net/http"
	"time"
)

func RenderFragmentTemplate(w http.ResponseWriter, html_file string) {
	time.Sleep(1 * time.Second)
	tmpl := template.Must(template.ParseFiles(html_file))
	tmpl.Execute(w, nil)
}
