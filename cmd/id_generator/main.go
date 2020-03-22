package main

import (
	"log"
	"mgoblin/users_grpc/internal/idgen"
	"net/http"
)

func main() {
	server := idgen.NewService()
	if err := http.ListenAndServe(":8080", server); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
