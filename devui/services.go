package devui

import "go.uber.org/zap"

type Services interface {
	GenerateDevUI() (devui string)
	ValidateDevUI(devui string) (valid bool)
	Backup()
	Restore()
}

func NewServices() Services {

	logger, _ := zap.NewProduction()

	devuis := make(map[string]bool, 0)

	return &service{
		logger.Sugar(),
		logger,
		devuis,
	}
}
