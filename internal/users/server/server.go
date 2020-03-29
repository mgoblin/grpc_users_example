package server

import (
	"context"
	"log"
	"os"
	"os/signal"

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
		URL:       config.IdGen.URL,
		Namespace: webservice.Namespace,
	}
	soapService := webservice.NewWsMath(&soapClient)

	server := &s.UserServiceServer{
		WebService: &soapService,
	}

	pbs.RegisterUserServiceServer(grpcServer, server)
	log.Printf("UserServiceServer registered")

	ctx := context.Background()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	return grpcServer
}
