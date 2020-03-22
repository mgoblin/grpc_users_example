package v1

import (
	"context"
	"reflect"
	"testing"

	pbs "mgoblin/users_grpc/internal/api/proto/v1"
	webservice "mgoblin/users_grpc/internal/api/wsdl/v1"
)

type WebServiceMock struct {
}

func (*WebServiceMock) Sum(Add *webservice.Add) (*webservice.Answer, error) {
	r := 42.0
	return &webservice.Answer{
		Result: &r,
	}, nil
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	var ws webservice.WsMath = &WebServiceMock{}
	service := UserServiceServer{WebService: &ws}
	req := pbs.ListUsersRequest{}
	want := &pbs.ListUsersResponse{
		Users: []*pbs.User{
			&pbs.User{Id: 42, FirstName: "Mike", LastName: "Downson"},
		},
	}

	res, err := service.ListUsers(ctx, &req)

	if err != nil {
		t.Errorf("Call list users has error %v", err)
	}

	if !reflect.DeepEqual(res.GetUsers(), want.GetUsers()) {
		t.Errorf("want: %+v, actual: %+v", *want, *res)
	}
}
