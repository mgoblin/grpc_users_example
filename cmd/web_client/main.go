package main

import (
	"log"
	"mgoblin/users_grpc/internal/web"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Users app web client starting")
	r := mux.NewRouter()
	h := handlers.LoggingHandler(os.Stdout, r)

	r.HandleFunc("/", web.IndexHandler())

	if err := http.ListenAndServe(":8081", h); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
