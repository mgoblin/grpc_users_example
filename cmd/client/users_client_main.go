package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	pbs "mgoblin/users_grpc/internal/api/proto/v1"
)

func main() {
	address := flag.String(
		"server",
		"localhost:7777",
		"gRPC server in format host:port")
	flag.Parse()

	log.Printf("Server address is %s", *address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbs.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersRequest := pbs.ListUsersRequest{PageSize: 20}

	usersListResponse, err := c.ListUsers(ctx, &usersRequest)

	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Printf("Create result: <%+v>\n\n", usersListResponse)
}
