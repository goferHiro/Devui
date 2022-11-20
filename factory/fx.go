package factory

import (
	"github.com/goferHiro/DevEUI/devui"
	"github.com/goferHiro/DevEUI/lorowan"
	"go.uber.org/zap"
)

func NewServices() Services {

	logger, _ := zap.NewDevelopment()
	var mq = make(chan string, 10)
	lorowanServices := lorowan.NewServices()
	devuiServices := devui.NewServices()

	devuiServices.Restore()

	return &service{
		logger.Sugar(),
		logger,
		mq,
		lorowanServices,
		devuiServices,
	}
}
