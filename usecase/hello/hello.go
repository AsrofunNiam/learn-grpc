package hello

import (
	v2 "github.com/AsrofunNiam/learn-grpc/proto/contracts/v2/contracts/v2"
	"github.com/AsrofunNiam/learn-grpc/repository"
	"gorm.io/gorm"
)

type Usecase interface {
	SayHello(req *v2.HelloRequest) *v2.HelloResponse
	SayHelloGateway(req *v2.HelloRequest) *v2.HelloResponse
}

type helloUsecase struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
}

func NewUsecase(
	productRepository repository.ProductRepository,
	db *gorm.DB,
) Usecase {
	return &helloUsecase{
		ProductRepository: productRepository,
		DB:                db,
	}
}

func (h *helloUsecase) SayHello(req *v2.HelloRequest) *v2.HelloResponse {
	return &v2.HelloResponse{
		Name:      req.Name,
		Age:       req.Age,
		Addresses: req.Addresses,
	}
}

func (h *helloUsecase) SayHelloGateway(req *v2.HelloRequest) *v2.HelloResponse {
	product := h.ProductRepository.FindByID(h.DB, 1)
	helloResponse := &v2.HelloResponse{}

	helloResponse.Name = req.Name
	helloResponse.Age = req.Age
	helloResponse.Addresses = req.Addresses

	if product.ID != 0 {
		helloResponse.Age += 10
	}

	return helloResponse
}
