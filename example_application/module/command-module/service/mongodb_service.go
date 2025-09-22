package service

import (
	"fmt"
	"github.com/lamxy/fiberhouse/example_application/module/command-module/model"
	"github.com/lamxy/fiberhouse/frame"
)

type MongodbService struct {
	*frame.Service
	MongoModel *model.MongodbModel
}

func NewMongodbService(ctx frame.ContextCommander, mongodbModel *model.MongodbModel) *MongodbService {
	return &MongodbService{
		Service:    frame.NewService(ctx).SetName("MongodbService").(*frame.Service),
		MongoModel: mongodbModel,
	}
}

func (s *MongodbService) Test() error {
	fmt.Println("MongodbService Test OK")
	return nil
}
