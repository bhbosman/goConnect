package stream

import (
	"context"
	"github.com/bhbosman/gocommon/comms/commsImpl"
	"go.uber.org/fx"
)

type ForexFactory struct {
	name   string
	apiKey string
}

func (self *ForexFactory) Create(name string, cancelCtx context.Context, cancelFunc context.CancelFunc, logger fx.ILogger) commsImpl.IConnectionReactor {
	result := newForex(logger, cancelCtx, cancelFunc, name, self.apiKey)
	return result
}

func (self *ForexFactory) Name() string {
	return self.name
}

func NewConnectionReactorForexFactory(name string, apiKey string) *ForexFactory {
	return &ForexFactory{name: name, apiKey: apiKey}
}
