package repository

import (
	"context"

	model "github.com/AsrofunNiam/learn-grpc/model/domain"
)

type UserRepository struct {
	// connection or db client
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (*model.User, error) {
	// Simulasikan pengambilan data dari database
	return &model.User{ID: id, Name: "John Doe", Email: "johndoe@example.com"}, nil
}
