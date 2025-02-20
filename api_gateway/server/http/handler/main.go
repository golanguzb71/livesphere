package handler

import (
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/server/grpc"
)

type Optoins struct {
	Service *grpc.GRPCService
	Logger  logger.Logger
	Cfg     *config.Config
}

type Handler struct {
	service *grpc.GRPCService
	logger  logger.Logger
	cfg     *config.Config
}

func NewHandler(options Optoins) *Handler {
	return &Handler{
		service: options.Service,
		logger:  options.Logger,
		cfg:     options.Cfg,
	}
}
