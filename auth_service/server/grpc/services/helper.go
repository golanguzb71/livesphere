package services

import (
	"github.com/golanguzb71/livesphere-auth-service/config"
	"github.com/golanguzb71/livesphere-auth-service/pkg/logger"
	"github.com/golanguzb71/livesphere-auth-service/server/grpc/client"
	"github.com/golanguzb71/livesphere-auth-service/storage"
)

type ServiceOptions struct {
	ServiceManager client.ServiceManager
	Storage        storage.StorageI
	Config         *config.Config
	Logger         logger.Logger
}
