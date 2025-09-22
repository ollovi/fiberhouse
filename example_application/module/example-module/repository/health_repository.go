package repository

import (
	"github.com/lamxy/fiberhouse/example_application/module/constant"
	"github.com/lamxy/fiberhouse/frame"
)

type HealthRepository struct {
	frame.RepositoryLocator
	Status string
}

func NewHealthRepository(ctx frame.ContextFramer) *HealthRepository {
	return &HealthRepository{
		RepositoryLocator: frame.NewRepository(ctx).SetName(GetKeyHealthRepository()),
		Status:            "Health is OK",
	}
}

func GetKeyHealthRepository(ns ...string) string {
	return frame.RegisterKeyName("HealthRepository", frame.GetNamespace([]string{constant.NameModuleExample}, ns...)...)
}

func RegisterKeyHealthRepository(ctx frame.ContextFramer, ns ...string) string {
	return frame.RegisterKeyInitializerFunc(GetKeyHealthRepository(ns...), func() (interface{}, error) {
		return NewHealthRepository(ctx), nil
	})
}

func (h *HealthRepository) Test() error {
	return nil
}
