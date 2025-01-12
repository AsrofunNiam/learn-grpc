package main

import (
	"context"
	"log"
	"time"

	"github.com/AsrofunNiam/learn-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Hubungkan ke server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Buat client
	client := proto.NewUserServiceClient(conn)

	// Kirim permintaan ke server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetUserByID(ctx, &proto.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	log.Printf("User: ID=%d, Name=%s, Email=%s", res.Id, res.Name, res.Email)
}
