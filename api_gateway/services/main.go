package services

import "github.com/golanguzb71/livesphere-api-gateway/config"

type ServiceManager interface {
	AuthService() AuthServiceI
}

type grpcClients struct {
	authService AuthServiceI
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	return &grpcClients{}, nil
}
