package worker

import (
	"context"
	"time"

	v2 "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"
	"github.com/AsrofunNiam/learn-grpc/usecase/hello"
	"gorm.io/gorm"
)

type HelloWorkerConfig struct {
	Interval time.Duration
}

type HelloWorker struct {
	Config  HelloWorkerConfig
	Usecase hello.Usecase
	Db      *gorm.DB
}

func NewHelloWorker(
	config HelloWorkerConfig,
	usecase hello.Usecase,
	db *gorm.DB,
) *HelloWorker {
	return &HelloWorker{
		Config:  config,
		Usecase: usecase,
		Db:      db,
	}
}

// Name mengembalikan nama worker
func (w *HelloWorker) Name() string {
	return "HelloWorker"
}

// HandleHelloRequest menangani request dan memprosesnya menggunakan usecase
func (w *HelloWorker) HandleHelloRequest(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	response := w.Usecase.SayHello(req)
	return response, nil
}

func (w *HelloWorker) HandleHelloGateway(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	response := w.Usecase.SayHelloGateway(req)
	return response, nil
}
