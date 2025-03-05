package services

import (
	"fmt"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type CoreServiceI interface {
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
