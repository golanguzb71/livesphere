package client

import "github.com/golanguzb71/livesphere-finance-service/config"

type ServiceManager interface {
}

type grpcClients struct {
}

// connect to external clients here
func NewGrpcClients(conf *config.Config) (ServiceManager, error) {

	return &grpcClients{}, nil
}
