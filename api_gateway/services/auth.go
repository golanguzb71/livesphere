package services

import (
	"fmt"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type AuthServiceI interface {
}
type AuthService struct {
	pool *grpcpool.Pool
	conn *grpc.ClientConn
}

func NewAuthService(cfg *config.Config) (AuthServiceI, error) {
	factory := func() (*grpc.ClientConn, error) {
		return grpc.NewClient(
			fmt.Sprintf("%s:%d", cfg.AuthServiceHost, cfg.AuthServicePort),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)))

	}

	pool, err := grpcpool.New(factory, 2, 4, time.Second, time.Minute*30)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.AuthServiceHost, cfg.AuthServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)))

	if err != nil {
		return nil, err
	}

	return &AuthService{
		pool: pool,
		conn: conn,
	}, nil
}
