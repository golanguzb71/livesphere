package repo

import pb "github.com/golanguzb71/livesphere-auth-service/genproto/template_service"

type HealthCheckRepo interface {
	HealthCheck() (*pb.HealthCheckResponse, error)
}
