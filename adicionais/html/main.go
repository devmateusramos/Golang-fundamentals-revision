package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		usuarioUm := usuario{
			"Mateus",
			"mateu@email.com",
		}

		templates.ExecuteTemplate(w, "home.html", usuarioUm)
	})

	fmt.Println("Server is running on port 5000")
	log.Fatalln(http.ListenAndServe(":5000", nil))
}
