package handler

import (
	"context"

	"github.com/AsrofunNiam/learn-grpc/proto"
	"github.com/AsrofunNiam/learn-grpc/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserByID menangani permintaan gRPC untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByID(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := h.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
