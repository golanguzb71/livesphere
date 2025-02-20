package storage

import (
	"github.com/golanguzb71/livesphere-api-gateway/config"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/db"
	"github.com/golanguzb71/livesphere-api-gateway/pkg/logger"
	"github.com/golanguzb71/livesphere-api-gateway/storage/postgres"
	"github.com/golanguzb71/livesphere-api-gateway/storage/repo"
)

type StorageI interface {
	HealthCheck() repo.HealthCheckRepo
}

type storagePg struct {
	healthCheck repo.HealthCheckRepo
}

func New(db *db.Postgres, log logger.Logger, cfg *config.Config) StorageI {
	return &storagePg{
		healthCheck: postgres.NewHealthCheckRepo(db, log, cfg),
	}
}

func (s *storagePg) HealthCheck() repo.HealthCheckRepo {
	return s.healthCheck
}
