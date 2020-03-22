package v1

import (
	"context"
	"log"

	pbs "mgoblin/users_grpc/internal/api/proto/v1"
	webservice "mgoblin/users_grpc/internal/api/wsdl/v1"
)

type UserServiceServer struct {
	WebService *webservice.WsMath
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pbs.ListUsersRequest) (*pbs.ListUsersResponse, error) {
	users := []*pbs.User{
		&pbs.User{Id: s.generateID(), FirstName: "Mike", LastName: "Downson"},
	}
	response := pbs.ListUsersResponse{
		Users: users,
	}

	return &response, nil
}

func (s *UserServiceServer) generateID() int64 {
	ws := *s.WebService
	x, y := 1.0, 2.0
	answer, err := ws.Sum(&webservice.Add{X: &x, Y: &y})
	log.Println("Call webservice")
	if err != nil {
		log.Println(err)
		return -1
	} else {
		return int64(*answer.Result)
	}
}
