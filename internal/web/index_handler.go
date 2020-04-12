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

func IndexHandler(client v1.UserServiceClient) func(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("%q", err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		users, err := client.ListUsers()

		data := &indexData{
			PageTitle: "Users",
			Users:     users,
			Error:     err,
		}
		tmpl.Execute(w, data)
	}
}
