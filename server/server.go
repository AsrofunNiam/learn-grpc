package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/AsrofunNiam/learn-grpc/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHelloBroh(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello, " + req.GetName()}, nil
}
func main() {
	// Mengatur listener pada port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Membuat server gRPC baru
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	// Menjalankan server
	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
