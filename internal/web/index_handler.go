package web

import (
	"log"
	"net/http"
	"text/template"

	pbs "mgoblin/users_grpc/internal/api/proto/v1"
	v1 "mgoblin/users_grpc/internal/service/v1/client"
)

type indexData struct {
	PageTitle string
	Users     []*pbs.User
	Error     error
}

func IndexHandler() func(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("%q", err)
	}

	address := "localhost:8080"

	c := v1.NewUsersClient(&address)

	return func(w http.ResponseWriter, req *http.Request) {
		users, err := c.ListUsers()

		data := &indexData{
			PageTitle: "title",
			Users:     users,
			Error:     err,
		}
		tmpl.Execute(w, data)
	}
}
