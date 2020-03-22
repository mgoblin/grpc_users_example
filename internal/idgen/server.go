package idgen

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewService() http.Handler {
	log.Println("Users webService server starting")
	r := mux.NewRouter()
	h := handlers.LoggingHandler(os.Stdout, r)

	r.HandleFunc("/", ListUsersHandler)

	return h
}
