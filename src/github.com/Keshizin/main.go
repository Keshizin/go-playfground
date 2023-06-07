package main

import (
	"html/template"
	"net/http"
)

type Login struct {
	User     string
	Password string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	logins := []Login{
		{User: "carlos1", Password: "123"},
		{User: "carlos2", Password: "456"},
		{User: "rodrigo", Password: "123"},
		{User: "marina", Password: "123"},
		{User: "jummy", Password: "123"},
		{User: "takeshi", Password: "123"},
	}

	templates.ExecuteTemplate(w, "Index", logins)
}
