package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AsrofunNiam/learn-grpc/hello" // Sesuaikan dengan path proyekmu

	"google.golang.org/grpc"
)

func main() {
	// Membuka koneksi ke server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Membuat client gRPC
	c := hello.NewGreeterClient(conn)

	// Membuat context dengan timeout 1 detik
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Memanggil RPC SayHello
	// r, err := c.SayHelloBroh(ctx, &hello.HelloRequest{Name: "testing"})
	r, err := c.SayHelloBroh(ctx, &hello.HelloRequest{Name: "testing"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Menampilkan hasil
	fmt.Println("Greeting:", r.GetMessage())
}
