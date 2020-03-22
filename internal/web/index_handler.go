package web

import (
	"log"
	"net/http"
	"text/template"
)

type indexData struct {
	PageTitle string
}

func IndexHandler() func(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("%q", err)
	}
	return func(w http.ResponseWriter, req *http.Request) {
		data := &indexData{PageTitle: "title"}
		tmpl.Execute(w, data)
	}
}
