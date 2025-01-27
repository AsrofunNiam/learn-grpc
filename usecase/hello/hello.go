package hello

import v2 "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"

type Usecase interface {
	SayHello(req *v2.HelloRequest) *v2.HelloResponse
	SayHelloGateway(req *v2.HelloRequest) *v2.HelloResponse
}

type helloUsecase struct{}

func NewUsecase() Usecase {
	return &helloUsecase{}
}

func (h *helloUsecase) SayHello(req *v2.HelloRequest) *v2.HelloResponse {
	return &v2.HelloResponse{
		Name:      req.Name,
		Age:       req.Age,
		Addresses: req.Addresses,
	}
}

func (h *helloUsecase) SayHelloGateway(req *v2.HelloRequest) *v2.HelloResponse {
	return &v2.HelloResponse{
		Name:      req.Name,
		Age:       req.Age,
		Addresses: req.Addresses,
	}
}
