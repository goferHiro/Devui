package factory

import (
	"github.com/goferHiro/DevEUI/lorowan"
	"go.uber.org/zap"
)

func NewServices() Services {

	logger, _ := zap.NewProduction()
	var mq = make(chan string, 100)
	lorowanServices := lorowan.NewServices()

	return &service{
		logger.Sugar(),
		logger,
		mq,
		lorowanServices,
	}
}
