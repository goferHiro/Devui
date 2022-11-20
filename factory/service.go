package factory

import (
	"github.com/goferHiro/DevEUI/lorowan"
	"go.uber.org/zap"
)

type service struct {
	log    *zap.SugaredLogger
	logger *zap.Logger
	mq     chan string

	lorowanServices lorowan.Services
}

func (s *service) Produce(devui string) {
	s.mq <- devui
	s.logger.Debug("produced", zap.String("devui", devui))
}

func (s *service) Consume() {
	devui := <-s.mq
	s.lorowanServices.RegisterDEVEUI(devui)
}
