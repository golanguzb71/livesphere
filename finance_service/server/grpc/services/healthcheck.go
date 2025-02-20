package services

import (
	"context"

	"github.com/golanguzb71/livesphere-finance-service/config"
	pb "github.com/golanguzb71/livesphere-finance-service/genproto/template_service"
	"github.com/golanguzb71/livesphere-finance-service/pkg/logger"
	"github.com/golanguzb71/livesphere-finance-service/server/grpc/client"
	"github.com/golanguzb71/livesphere-finance-service/storage"
)

type HealthCheckService struct {
	storage  storage.StorageI
	logger   logger.Logger
	cfg      *config.Config
	services client.ServiceManager
	pb.UnimplementedHealthCheckServiceServer
}

func NewHealthCheckService(options *ServiceOptions) pb.HealthCheckServiceServer {
	return &HealthCheckService{
		storage:  options.Storage,
		logger:   options.Logger,
		services: options.ServiceManager,
		cfg:      options.Config,
	}
}

func (s *HealthCheckService) HealthCheck(ctx context.Context, req *pb.Empty) (*pb.HealthCheckResponse, error) {
	return s.storage.HealthCheck().HealthCheck()
}
