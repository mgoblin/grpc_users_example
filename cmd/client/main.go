package main

import (
	"log"

	v1 "mgoblin/users_grpc/internal/service/v1/client"
	config "mgoblin/users_grpc/internal/users/client/config"
)

func main() {

	cfg := config.New()
	address := &cfg.UserService.URL

	log.Printf("Server address is %s", *address)

	c := v1.NewUsersClient(address)
	users, err := c.ListUsers()

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Printf("Create result: <%+v>\n\n", users)
	c.Close()
}
