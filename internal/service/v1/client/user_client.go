package client

import (
	"context"
	"log"
	pbs "mgoblin/users_grpc/internal/api/proto/v1"
	"time"

	"google.golang.org/grpc"
)

type UserServiceClient struct {
	client     pbs.UserServiceClient
	connection *grpc.ClientConn
}

func NewUsersClient(address *string) *UserServiceClient {

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	gprcClient := pbs.NewUserServiceClient(conn)

	return &UserServiceClient{client: gprcClient, connection: conn}
}

func (s *UserServiceClient) Close() {
	s.connection.Close()
}

func (s *UserServiceClient) ListUsers() ([]*pbs.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersRequest := pbs.ListUsersRequest{PageSize: 20}

	usersListResponse, err := s.client.ListUsers(ctx, &usersRequest)

	if err != nil {
		return nil, err
	}

	return usersListResponse.Users, nil
}
