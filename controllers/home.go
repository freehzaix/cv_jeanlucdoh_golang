package controllers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/home.html"))
    tmpl.Execute(w, nil)
}