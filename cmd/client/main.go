package main

import (
	"flag"
	"log"

	v1 "mgoblin/users_grpc/internal/service/v1/client"
)

func main() {
	address := flag.String(
		"server",
		"localhost:7777",
		"gRPC server in format host:port")
	flag.Parse()

	log.Printf("Server address is %s", *address)

	c := v1.NewUsersClient(address)
	users, err := c.ListUsers()

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Printf("Create result: <%+v>\n\n", users)
}
