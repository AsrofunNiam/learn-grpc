package main

import (
	"log"
	"net"
	"net/http"

	"github.com/AsrofunNiam/learn-grpc/app"
	c "github.com/AsrofunNiam/learn-grpc/configuration"
	"github.com/AsrofunNiam/learn-grpc/gapi"
	"github.com/AsrofunNiam/learn-grpc/helper"
	"github.com/AsrofunNiam/learn-grpc/pb"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load configuration:", err)
	}

	// Connect to database
	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)

	// Run gRPC server in a separate goroutine
	go func() {
		runGrpcServer(configuration, db)
	}()

	// Initialize and run HTTP server
	validate := validator.New()
	router := app.NewRouter(db, validate)
	port := configuration.HttpPort

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("HTTP server is running on port %s", port)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

func runGrpcServer(config c.Configuration, store *gorm.DB) {
	// Initialize gRPC server
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create gRPC server:", err)
	}

	grpcServer := grpc.NewServer()

	// Register services and reflection
	pb.RegisterSimpleBankServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	// Start listening on gRPC port
	listener, err := net.Listen("tcp", ":"+config.GRPCPort)

	if err != nil {
		log.Fatal("Cannot create gRPC listener:", err)
	}

	log.Printf("gRPC server is running on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start gRPC server:", err)
	}
}
