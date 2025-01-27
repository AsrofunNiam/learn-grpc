package worker

import (
	"context"
	"time"

	v2 "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"
	"github.com/AsrofunNiam/learn-grpc/usecase/hello"
)

type HelloWorkerConfig struct {
	Interval time.Duration
}

type HelloWorker struct {
	config  HelloWorkerConfig
	usecase hello.Usecase
}

func NewHelloWorker(config HelloWorkerConfig, usecase hello.Usecase) *HelloWorker {
	if config.Interval == 0 {
		config.Interval = 5 * time.Second // Default interval jika tidak disetel
	}

	return &HelloWorker{
		config:  config,
		usecase: usecase,
	}
}

// Name mengembalikan nama worker
func (w *HelloWorker) Name() string {
	return "HelloWorker"
}

// HandleHelloRequest menangani request dan memprosesnya menggunakan usecase
func (w *HelloWorker) HandleHelloRequest(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	response := w.usecase.SayHello(req)
	return response, nil
}

func (w *HelloWorker) HandleHelloGateway(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	response := w.usecase.SayHello(req)
	return response, nil
}
