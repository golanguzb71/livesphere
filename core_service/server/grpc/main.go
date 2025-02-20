package grpc

import (
	"fmt"
	"net"

	"github.com/golanguzb71/livesphere-core-service/config"
	pb "github.com/golanguzb71/livesphere-core-service/genproto/template_service"
	"github.com/golanguzb71/livesphere-core-service/pkg/db"
	"github.com/golanguzb71/livesphere-core-service/pkg/logger"
	"github.com/golanguzb71/livesphere-core-service/server/grpc/client"
	"github.com/golanguzb71/livesphere-core-service/server/grpc/services"
	"github.com/golanguzb71/livesphere-core-service/storage"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

type GRPCService struct {
	HealthCheckService pb.HealthCheckServiceServer
}

func New(cfg *config.Config, log logger.Logger) (*GRPCService, error) {
	psql, err := db.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	storageObj := storage.New(psql, log, cfg)

	grpcClient, err := client.NewGrpcClients(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while connecting with grpc clients: %v", err)
	}

	serviceOptions := &services.ServiceOptions{
		ServiceManager: grpcClient,
		Storage:        storageObj,
		Config:         cfg,
		Logger:         log,
	}

	return &GRPCService{
		HealthCheckService: services.NewHealthCheckService(serviceOptions),
	}, nil
}

func (service *GRPCService) Run(logger logger.Logger, cfg *config.Config) {
	server := grpc.NewServer()

	pb.RegisterHealthCheckServiceServer(server, service.HealthCheckService)

	listener, err := net.Listen("tcp", ":"+cfg.RPCPort)
	if err != nil {
		panic("Error while creating listener")
	}

	logger.Info("GRPC server is running on port " + cfg.RPCPort)

	err = server.Serve(listener)
	if err != nil {
		panic("error while serving gRPC server on port " + cfg.RPCPort)
	}
}
