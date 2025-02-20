package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/server/grpc"
	"github.com/golanguzb71/livesphere-api-gateway/server/http/handler"
)

type HttpServerI interface {
	Run() error
}

type httpServer struct {
	handler *handler.Handler
}

func New(cfg *config.Config, log logger.Logger, grpc *grpc.GRPCService) (*httpServer, error) {
	options := handler.Optoins{
		Logger:  log,
		Cfg:     cfg,
		Service: grpc,
	}
	handler := handler.NewHandler(options)

	return &httpServer{
		handler: handler,
	}, nil
}

func (s *httpServer) Run(log logger.Logger, cfg *config.Config) error {
	log.Info("Starting http server", logger.String("port", ":"+cfg.HTTPPort))

	engine := gin.Default()
	engine.Use(gin.Recovery())

	engine.GET("/healthcheck", s.handler.HealthCheck)

	err := engine.Run(":" + cfg.HTTPPort)

	return err
}
