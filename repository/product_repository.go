package repository

import (
	"context"

	model "github.com/AsrofunNiam/learn-grpc/model/domain"
)

type ProductRepository struct {
	// connection or db client
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int32) (*model.Product, error) {
	// Simulasikan pengambilan data dari database
	return &model.Product{ID: id, Name: "Laptop", Price: 1000}, nil
}
