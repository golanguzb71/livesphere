package services

import (
	"github.com/golanguzb71/livesphere-core-service/config"
	"github.com/golanguzb71/livesphere-core-service/pkg/logger"
	"github.com/golanguzb71/livesphere-core-service/server/grpc/client"
	"github.com/golanguzb71/livesphere-core-service/storage"
)

type ServiceOptions struct {
	ServiceManager client.ServiceManager
	Storage        storage.StorageI
	Config         *config.Config
	Logger         logger.Logger
}
