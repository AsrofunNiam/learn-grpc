package main

import (
	"context"
	"fmt"
	"log"
	"time"

	v2 "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Koneksi ke server gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Client Greeter
	client := v2.NewGreeterClient(conn)

	// Request value
	request := &v2.HelloRequest{
		Name: "Alice",
		Age:  25,
		Addresses: []*v2.Address{
			{Street: "Jalan Raya", City: "Bandung", Country: "Indonesia"},
			{Street: "Jalan Jakarta", City: "Jakarta", Country: "Indonesia"},
		},
	}

	// Set timeout request context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send request to server
	response, err := client.SayHelloBroh(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Response result
	fmt.Printf("Response from server: Name: %s, Age: %d, Address: %v\n",
		response.GetName(), response.GetAge(), response.GetAddresses())
}
