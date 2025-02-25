package main

import (
	rediscache "github.com/golanguzb70/redis-cache"
	"github.com/golanguzb71/livesphere-api-gateway/api"
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "livesphere_go_api_gateway")
	gprcClients, err := services.NewGrpcClients(&cfg)
	if err != nil {
		log.Error("Error while getting grpc clients", logger.Error(err))
		return
	}

	cache := rediscache.New(&rediscache.Config{
		RedisHost: cfg.RedisHost,
		RedisPort: cfg.RedisPort,
	})

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      &cfg,
		Services: gprcClients,
		Cache:    cache,
	})

	server.Run(":" + cfg.HttpPort)
}
