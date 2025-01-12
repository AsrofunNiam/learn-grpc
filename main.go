package main

import (
	"context"
	"log"
	"net"

	"github.com/AsrofunNiam/learn-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct mengimplementasikan interface proto.UserServiceServer
type Server struct {
	proto.UnimplementedUserServiceServer
}

// GetUserByID adalah implementasi dari RPC
func (s *Server) GetUserByID(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	// Dummy data
	userData := map[int32]*proto.GetUserResponse{
		1: {Id: 1, Name: "John Doe", Email: "john.doe@example.com"},
		2: {Id: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
	}

	// Cari user berdasarkan ID
	if user, found := userData[req.Id]; found {
		return user, nil
	}

	// Jika tidak ditemukan, kembalikan error
	return nil, grpc.Errorf(grpc.Code(grpc.ErrClientConnTimeout), "User not found")
}

func main() {
	// Listener untuk server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Inisialisasi gRPC server
	grpcServer := grpc.NewServer()

	// Registrasi service
	proto.RegisterUserServiceServer(grpcServer, &Server{})

	// Reflection untuk debugging dengan tools seperti grpcurl
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
