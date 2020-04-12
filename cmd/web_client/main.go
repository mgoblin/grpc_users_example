package main

import (
	"fmt"
	"log"
	"mgoblin/users_grpc/internal/web"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	v1 "mgoblin/users_grpc/internal/service/v1/client"
	config "mgoblin/users_grpc/internal/users/client/config"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	log.Println("Users app web client starting")
	r := mux.NewRouter()
	h := handlers.LoggingHandler(os.Stdout, r)

	cfg := config.New()
	usersAddress := cfg.UserService.URL
	port := cfg.WebConfig.TCPPort

	c := v1.NewUsersClient(&usersAddress)
	defer c.Close()

	log.Printf("Listening on %s", fmt.Sprintf(":%d", port))
	log.Printf("Users service address is %s", usersAddress)

	r.HandleFunc("/", web.IndexHandler(*c))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), h); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
