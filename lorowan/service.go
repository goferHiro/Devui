package lorowan

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type service struct {
	host   string
	log    *zap.SugaredLogger
	logger *zap.Logger
}

func (s *service) RegisterDEVEUI(devui string) (err error) {

	var payload struct {
		name string
	}

	payload.name = devui

	url := fmt.Sprintf("%s/sensor-onboarding-sample", s.host)

	_, body, errs := gorequest.New().Post(url).Send(payload).End()

	if body == "OK" {
		s.log.Info("registered DEVUI-", devui)
		return
	}

	if errs == nil {
		err = fmt.Errorf("failed to receive OK")
		s.logger.Error("didn't receive the right response from register provider", zap.Error(err))
	} else {
		err = multierr.Combine(errs...)
	}

	s.logger.Error("registration of new factory failed due to ", zap.String("factory", devui), zap.Errors("errors", errs))

	return

}
