package handler

import (
	"context"

	"github.com/AsrofunNiam/learn-grpc/proto"
	"github.com/AsrofunNiam/learn-grpc/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetProductByID menangani permintaan gRPC untuk mendapatkan produk berdasarkan ID
func (h *ProductHandler) GetProductByID(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	product, err := h.productService.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.GetProductResponse{
		Id:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
