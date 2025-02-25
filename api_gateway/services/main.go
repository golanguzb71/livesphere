package services

import "github.com/golanguzb71/livesphere-api-gateway/config"

type ServiceManager interface {
}

type grpcClients struct {
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	return &grpcClients{}, nil
}
