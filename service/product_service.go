package service

import (
	"context"

	model "github.com/AsrofunNiam/learn-grpc/model/domain"
	"github.com/AsrofunNiam/learn-grpc/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

// GetProduct mengembalikan informasi Product berdasarkan ID
func (s *ProductService) GetProduct(ctx context.Context, id int32) (*model.Product, error) {
	return s.productRepo.GetProductByID(ctx, id)
}
