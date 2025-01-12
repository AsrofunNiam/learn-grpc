package gapi

import (
	c "github.com/AsrofunNiam/learn-grpc/configuration"
	"github.com/AsrofunNiam/learn-grpc/pb"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedSimpleBankServiceServer
	config c.Configuration
	store  *gorm.DB
}

// NewServer creates a new gRPC server and sets up routes
func NewServer(config c.Configuration, store *gorm.DB) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
