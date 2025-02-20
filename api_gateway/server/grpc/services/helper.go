package services

import (
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/server/grpc/client"
	"github.com/golanguzb71/livesphere-api-gateway/storage"
)

type ServiceOptions struct {
	ServiceManager client.ServiceManager
	Storage        storage.StorageI
	Config         *config.Config
	Logger         logger.Logger
}
