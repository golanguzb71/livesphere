package services

import (
	"github.com/golanguzb71/livesphere-api-gateway/config"
)

type ServiceManager interface {
	AuthService() AuthServiceI
	CoreService() CoreServiceI
	PaymentService() PaymentServiceI
}

type grpcClients struct {
	authService    AuthServiceI
	coreService    CoreServiceI
	paymentService PaymentServiceI
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	authService, err := NewAuthService(conf)
	if err != nil {
		return nil, err
	}
	coreService, err := NewCoreService(conf)
	if err != nil {
		return nil, err
	}
	paymentService, err := NewPaymentService(conf)
	if err != nil {
		return nil, err
	}
	return &grpcClients{
		authService:    authService,
		coreService:    coreService,
		paymentService: paymentService,
	}, nil
}
func (g *grpcClients) AuthService() AuthServiceI {
	return g.authService
}
func (g *grpcClients) CoreService() CoreServiceI {
	return g.coreService
}
func (g *grpcClients) PaymentService() PaymentServiceI {
	return g.paymentService
}
