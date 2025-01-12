package service

import (
	"context"

	model "github.com/AsrofunNiam/learn-grpc/model/domain"
	"github.com/AsrofunNiam/learn-grpc/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetUser mengembalikan informasi User berdasarkan ID
func (s *UserService) GetUser(ctx context.Context, id int32) (*model.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}
