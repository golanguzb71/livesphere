package services

import (
	"github.com/golanguzb71/livesphere-finance-service/config"
	"github.com/golanguzb71/livesphere-finance-service/pkg/logger"
	"github.com/golanguzb71/livesphere-finance-service/server/grpc/client"
	"github.com/golanguzb71/livesphere-finance-service/storage"
)

type ServiceOptions struct {
	ServiceManager client.ServiceManager
	Storage        storage.StorageI
	Config         *config.Config
	Logger         logger.Logger
}
