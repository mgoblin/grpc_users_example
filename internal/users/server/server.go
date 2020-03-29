package server

import (
	"log"

	"github.com/fiorix/wsdl2go/soap"
	"google.golang.org/grpc"

	pbs "mgoblin/users_grpc/internal/api/proto/v1"
	webservice "mgoblin/users_grpc/internal/api/wsdl/v1"
	s "mgoblin/users_grpc/internal/service/v1"
	cfg "mgoblin/users_grpc/internal/users/server/config"
)

func New(config *cfg.Config) *grpc.Server {
	grpcServer := grpc.NewServer()
	log.Printf("gRPC server created")

	soapClient := soap.Client{
		URL:       "http://127.0.0.1:8080",
		Namespace: webservice.Namespace,
	}
	soapService := webservice.NewWsMath(&soapClient)

	server := &s.UserServiceServer{
		WebService: &soapService,
	}

	pbs.RegisterUserServiceServer(grpcServer, server)
	log.Printf("UserServiceServer registered")

	return grpcServer
}
