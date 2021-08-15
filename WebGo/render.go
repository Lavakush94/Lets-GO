package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}

}
