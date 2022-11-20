package lorowan

import "go.uber.org/zap"

type Services interface {
	RegisterDEVEUI(devui string) (err error)
}

func NewServices() Services {
	logger, _ := zap.NewProduction()
	return &service{
		"https://europe-west1-machinemax-dev-d524.cloudfunctions.net",
		logger.Sugar(),
		logger,
	}
}
