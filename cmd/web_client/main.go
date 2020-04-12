package main

import (
	"log"
	"mgoblin/users_grpc/internal/web"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	v1 "mgoblin/users_grpc/internal/service/v1/client"
)

func main() {
	log.Println("Users app web client starting")
	r := mux.NewRouter()
	h := handlers.LoggingHandler(os.Stdout, r)

	address := "localhost:7777"

	c := v1.NewUsersClient(&address)
	defer c.Close()

	r.HandleFunc("/", web.IndexHandler(*c))

	if err := http.ListenAndServe(":8081", h); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
