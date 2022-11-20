package factory

import (
	"github.com/goferHiro/DevEUI/devui"
	"github.com/goferHiro/DevEUI/lorowan"
	"go.uber.org/zap"
	"os"
)

func NewServices() Services {

	logger, _ := zap.NewDevelopment()

	if os.Getenv("mode") == "production" {
		logger, _ = zap.NewProduction()
	}

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
