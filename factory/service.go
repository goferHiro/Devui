package factory

import (
	"github.com/goferHiro/DevEUI/devui"
	"github.com/goferHiro/DevEUI/lorowan"
	"go.uber.org/zap"
	"sync"
)

type service struct {
	log    *zap.SugaredLogger
	logger *zap.Logger
	mq     chan string

	lorowanServices lorowan.Services
	devuiServices   devui.Services
}

func (s *service) Produce(devui string) {
	s.mq <- devui
	s.logger.Debug("produced", zap.String("devui", devui))
}

func (s *service) Consume() {
	devui := <-s.mq
	s.lorowanServices.RegisterDEVEUI(devui)
	s.logger.Debug("consumed", zap.String("devui", devui))

}

func (s *service) BatchOf100() (devuis []string) {
	devuis = make([]string, 0)
	for len(devuis) < 100 {
		devui := s.devuiServices.GenerateDevUI()
		if !s.devuiServices.ValidateDevUI(devui) {
			devuis = append(devuis, devui)
		}
	}

	return
}

func (s *service) ProduceBatch100(devuis []string) {
	var wg sync.WaitGroup
	for _, devui := range devuis {
		s.Produce(devui)

		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Consume()
		}()
	}
	wg.Wait()
	s.log.Info("backing up")
	s.devuiServices.Backup()
	s.log.Info("backup complete")
}
