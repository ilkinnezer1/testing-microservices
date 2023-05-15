package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 5000")
	err := http.ListenAndServe(":5000", nil)
	HandleError(err)
}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"../web/templates/base.layout.gohtml",
		"../web/templates/header.partial.gohtml",
		"../web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("../web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
