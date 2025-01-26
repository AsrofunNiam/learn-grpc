package main

import (
	"context"
	"log"
	"net"

	"github.com/AsrofunNiam/learn-grpc/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct
type server struct {
	hello.UnimplementedGreeterServer
}

// Implementasi SayHelloBroh
func (s *server) SayHelloBroh(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	response := &hello.HelloResponse{
		Name:      req.GetName(),
		Age:       req.GetAge(),
		Addresses: req.GetAddresses(),
	}
	return response, nil
}

func main() {
	// Listen di port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// GRPC server connection
	grpcServer := grpc.NewServer()

	// Register Greeter service
	hello.RegisterGreeterServer(grpcServer, &server{})

	// Register reflection client on other tools ( postman, other grpc client )
	reflection.Register(grpcServer)

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
