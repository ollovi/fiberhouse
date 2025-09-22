package service

import (
	"github.com/lamxy/fiberhouse/example_application/module/constant"
	"github.com/lamxy/fiberhouse/example_application/module/example-module/repository"
	"github.com/lamxy/fiberhouse/frame"
)

type HealthService struct {
	frame.ServiceLocator
	Resp *repository.HealthRepository
}

func NewHealthService(ctx frame.ContextFramer, resp *repository.HealthRepository) *HealthService {
	name := GetKeyHealthService()
	return &HealthService{
		ServiceLocator: frame.NewService(ctx).SetName(name),
		Resp:           resp,
	}
}

func GetKeyHealthService(ns ...string) string {
	return frame.RegisterKeyName("HealthService", frame.GetNamespace([]string{constant.NameModuleExample}, ns...)...)
}

func (s *HealthService) GetHealth() string {
	return s.Resp.Status
}
