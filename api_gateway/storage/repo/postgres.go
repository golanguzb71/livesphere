package repo

import pb "github.com/golanguzb71/livesphere-api-gateway/genproto/template_service"

type HealthCheckRepo interface {
	HealthCheck() (*pb.HealthCheckResponse, error)
}
