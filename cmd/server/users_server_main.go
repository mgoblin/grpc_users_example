package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fiorix/wsdl2go/soap"
	"google.golang.org/grpc"

	pbs "mgoblin/users_grpc/internal/api/proto/v1"
	webservice "mgoblin/users_grpc/internal/api/wsdl/v1"
	s "mgoblin/users_grpc/internal/service/v1"
)

func main() {
	port := 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	log.Printf("Listening port %d", port)

	// create a gRPC server object
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

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
