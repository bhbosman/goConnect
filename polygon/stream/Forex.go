package stream

import (
	"bytes"
	"context"
	"fmt"
	rxgo "github.com/ReactiveX/RxGo"
	polygonMessage "github.com/bhbosman/goMessages/polygon/stream"
	"github.com/bhbosman/gocommon/comms/connectionManager"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/multiBlock"
	"github.com/bhbosman/gocommon/stacks/websocket/wsmsg"
	"github.com/bhbosman/gocommon/stream"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/fx"
	"net"
	"net/url"
)

type Forex struct {
	PolygonConnectionReactor
	Status        string
	messageRouter *messageRouter.MessageRouter
}

func (self *Forex) Init(
	conn net.Conn,
	url *url.URL,
	connectionId string,
	connectionManager connectionManager.IConnectionManager,
	toConnectionFunc stream.ToConnectionFunc,
	toConnectionReactor stream.ToReactorFunc) (rxgo.NextExternalFunc, error) {
	_, _ = self.BaseConnectionReactor.Init(conn, url, connectionId, connectionManager, toConnectionFunc, toConnectionReactor)
	_ = self.messageRouter.Add(self.HandleReaderWriter)
	_ = self.messageRouter.Add(self.HandleWebSocketMessage)
	_ = self.messageRouter.Add(self.HandlePolygonForexCombined)

	return self.doNext, nil
}

func (self *Forex) HandleReaderWriter(msg *multiBlock.ReaderWriter) error {
	marshal, err := stream.UnMarshal(msg, self.CancelCtx, self.CancelFunc, self.ToReactor, self.ToConnection)
	if err != nil {
		return err
	}
	_, err = self.messageRouter.Route(marshal)
	return err
}

func (self *Forex) HandlePolygonForexCombined(msg *polygonMessage.PolygonForexCombined) error {
	switch msg.EventType {
	case "status":
		self.Status = msg.StatusStatus
		self.Logger.Printf(fmt.Sprintf("Status: %v, Message: %v", msg.StatusStatus, msg.StatusMessage))
		if self.Status == "connected" {
			authMessage := &polygonMessage.PolygonMessageSend{
				Action: "auth",
				Params: self.ApiKey,
			}
			self.SendMessage(authMessage)
		}
	}
	return nil
}

func (self *Forex) HandleWebSocketMessage(msg *wsmsg.WebSocketMessageWrapper) error {
	switch msg.Data.OpCode {
	case wsmsg.WebSocketMessage_OpText:
		m := &polygonMessage.PolygonForexCombined{}
		err := jsonpb.Unmarshal(bytes.NewBuffer(msg.Data.Message[1:len(msg.Data.Message)-1]), m)
		if err != nil {
			return err
		}
		err = self.ToReactor(true, m)
		if err != nil {
			return err
		}
		return nil
	case wsmsg.WebSocketMessage_OpStartLoop:
		return nil
	default:
		return nil
	}
}

func (self *Forex) doNext(external bool, i interface{}) {
	_, err := self.messageRouter.Route(i)
	if err != nil {
		return
	}
}

func (self *Forex) SendMessage(message proto.Message) error {
	rws := multiBlock.NewReaderWriter()
	m := jsonpb.Marshaler{}
	err := m.Marshal(rws, message)
	if err != nil {
		return err
	}
	flatten, err := rws.Flatten()
	if err != nil {
		return err
	}
	WebSocketMessage := wsmsg.WebSocketMessage{
		OpCode:  wsmsg.WebSocketMessage_OpText,
		Message: flatten,
	}
	readWriterSize, err := stream.Marshall(&WebSocketMessage)
	if err != nil {
		return err
	}

	return self.ToConnection(readWriterSize)
}

func newForex(
	logger fx.ILogger,
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	name string,
	apiKey string) *Forex {
	result := &Forex{
		PolygonConnectionReactor: NewPolygonConnectionReactor(logger,
			name,
			cancelCtx,
			cancelFunc,
			apiKey),
		Status:        "",
		messageRouter: messageRouter.NewMessageRouter(),
	}
	return result
}
