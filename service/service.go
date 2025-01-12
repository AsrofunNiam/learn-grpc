package service

// import (
// 	"context"

// 	model "github.com/AsrofunNiam/learn-grpc/model/domain"
// 	"github.com/AsrofunNiam/learn-grpc/repository"
// )

// type UserService struct {
// 	userRepo *repository.UserRepository
// }

// func NewUserService(userRepo *repository.UserRepository) *UserService {
// 	return &UserService{userRepo: userRepo}
// }

// func (s *UserService) GetUser(ctx context.Context, id int32) (*model.User, error) {
// 	return s.userRepo.GetUserByID(ctx, id)
// }

// type ProductService struct {
// 	productRepo *repository.ProductRepository
// }

// func NewProductService(productRepo *repository.ProductRepository) *ProductService {
// 	return &ProductService{productRepo: productRepo}
// }

// func (s *ProductService) GetProduct(ctx context.Context, id int32) (*model.Product, error) {
// 	return s.productRepo.GetProductByID(ctx, id)
// }
