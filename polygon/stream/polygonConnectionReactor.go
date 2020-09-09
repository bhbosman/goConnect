package stream

import (
	"context"
	"github.com/bhbosman/gocommon/comms/commsImpl"
	"go.uber.org/fx"
)

type PolygonConnectionReactor struct {
	commsImpl.BaseConnectionReactor
	ApiKey string
}

func NewPolygonConnectionReactor(
	logger fx.ILogger,
	name string,
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	apiKey string) PolygonConnectionReactor {
	return PolygonConnectionReactor{
		BaseConnectionReactor: commsImpl.NewBaseConnectionReactor(logger, name, cancelCtx, cancelFunc),
		ApiKey:                apiKey,
	}
}
