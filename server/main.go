package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	service "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"
)

type server struct {
	service.UnimplementedGreeterServer
}

// Implementasi metode SayHelloGrpcGateway
func (s *server) SayHelloGrpcGateway(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		Name:      req.Name,
		Age:       req.Age,
		Addresses: req.Addresses,
	}, nil
}

// Implementasi SayHelloBroh
func (s *server) SayHelloBroh(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	response := &service.HelloResponse{
		Name:      req.GetName(),
		Age:       req.GetAge(),
		Addresses: req.GetAddresses(),
	}
	return response, nil
}

func main() {
	// Start gRPC server
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		service.RegisterGreeterServer(grpcServer, &server{})

		log.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server for gRPC-Gateway
	mux := runtime.NewServeMux() // Inisialisasi multiplexer HTTP-Gateway

	// Register gRPC-Gateway handler
	err := service.RegisterGreeterHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:50051", // Alamat server gRPC
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Println("HTTP server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
