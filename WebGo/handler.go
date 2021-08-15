package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}

}
