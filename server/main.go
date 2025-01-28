package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/AsrofunNiam/learn-grpc/configuration"
	"github.com/AsrofunNiam/learn-grpc/database"
	v2 "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"
	"github.com/AsrofunNiam/learn-grpc/repository"
	"github.com/AsrofunNiam/learn-grpc/usecase/hello"
	"github.com/AsrofunNiam/learn-grpc/worker"
)

type server struct {
	v2.UnimplementedGreeterServer
	worker *worker.HelloWorker
}

func (s *server) SayHelloGrpcGateway(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	return s.worker.HandleHelloGateway(ctx, req)
}

func (s *server) SayHelloBroh(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	return s.worker.HandleHelloRequest(ctx, req)
}

func main() {

	configuration, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := database.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)

	// Inisialisasi usecase
	productRepository := repository.NewProductRepository()
	helloUsecase := hello.NewUsecase(productRepository, db)
	helloWorker := worker.NewHelloWorker(worker.HelloWorkerConfig{}, helloUsecase, db)

	// Start gRPC server
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		v2.RegisterGreeterServer(grpcServer, &server{worker: helloWorker})

		log.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server for gRPC-Gateway
	mux := runtime.NewServeMux()

	err = v2.RegisterGreeterHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:50051",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Println("HTTP server is running on port:", configuration.Port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
