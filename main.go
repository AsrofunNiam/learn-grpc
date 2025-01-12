package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AsrofunNiam/learn-grpc/handler"
	"github.com/AsrofunNiam/learn-grpc/proto"
	"github.com/AsrofunNiam/learn-grpc/repository"
	"github.com/AsrofunNiam/learn-grpc/service"
	"google.golang.org/grpc"
)

// Server struct mengimplementasikan interface proto.UserServiceServer
type ServerUser struct {
	proto.UnimplementedUserServiceServer
}
type ServerProduct struct {
	proto.UnimplementedProductServiceServer
}

func main() {
	// Initialize repositories
	userRepo := repository.NewUserRepository()
	productRepo := repository.NewProductRepository()

	// Initialize services
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	// Set up the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register services
	proto.RegisterUserServiceServer(grpcServer, &ServerUser{})
	proto.RegisterProductServiceServer(grpcServer, &ServerProduct{})

	fmt.Println("gRPC server is running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
