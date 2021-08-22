package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}

}

func RenderTemplateTest(w http.ResponseWriter, tmpl string) (map[string]*template.Template{}, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return err
		}

		matches, err := filepath.Glob("./*.layout.html")

		if err != nil {
			return err
		}

		if len(matches) > 0 {
			_, err := ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return err
			}
		}
		myCache[name] = ts
	}
	return nil
}
