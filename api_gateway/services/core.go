package services

import (
	"context"
	"fmt"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	coreService "github.com/golanguzb71/livesphere-api-gateway/genproto/core_service"
	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type CoreServiceI interface {
	LeadService(ctx context.Context) coreService.LeadServiceClient
	ExpectService(ctx context.Context) coreService.ExpectServiceClient
	SetService(ctx context.Context) coreService.SetServiceClient
	LeadDataService(ctx context.Context) coreService.LeadDataServiceClient
}
type CoreService struct {
	pool *grpcpool.Pool
	conn *grpc.ClientConn
}

func NewCoreService(cfg *config.Config) (CoreServiceI, error) {
	factory := func() (*grpc.ClientConn, error) {
		return grpc.NewClient(
			fmt.Sprintf("%s:%d", cfg.CoreServiceHost, cfg.CoreServicePort),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)))

	}

	pool, err := grpcpool.New(factory, 2, 4, time.Second, time.Minute*30)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.CoreServiceHost, cfg.CoreServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)))

	if err != nil {
		return nil, err
	}

	return &CoreService{
		pool: pool,
		conn: conn,
	}, nil
}

func (c *CoreService) LeadService(ctx context.Context) coreService.LeadServiceClient {
	return coreService.NewLeadServiceClient(c.conn)
}

func (c *CoreService) ExpectService(ctx context.Context) coreService.ExpectServiceClient {
	return coreService.NewExpectServiceClient(c.conn)
}

func (c *CoreService) SetService(ctx context.Context) coreService.SetServiceClient {
	return coreService.NewSetServiceClient(c.conn)
}

func (c *CoreService) LeadDataService(ctx context.Context) coreService.LeadDataServiceClient {
	return coreService.NewLeadDataServiceClient(c.conn)
}
