package internal

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/goMessages/polygon/rest"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io"
	"net/http"
	"strings"
)

func ReadJsonArray(body io.Reader) ([]json.RawMessage, error) {
	var rawDataForItems []json.RawMessage
	decoder := json.NewDecoder(body)

	err := decoder.Decode(&rawDataForItems)
	if err != nil {
		return nil, err
	}
	return rawDataForItems, nil
}

func ReadJsonError(body io.Reader, statusCode int) error {
	unmarshaler := &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
		AnyResolver:        nil,
	}
	errorResponse := rest.GetErrorResponse{}
	err := unmarshaler.Unmarshal(body, &errorResponse)
	if err != nil {
		return err
	}
	return NewError(statusCode, errorResponse.ErrorMessage, errorResponse.RequestId)
}

func ReadCommonError(statusCode int) error {
	switch {
	case statusCode == http.StatusUnauthorized:
		return NewError(statusCode, "Unauthorized - Check our API Key and account status", "")
	case statusCode == http.StatusNotFound:
		return NewError(statusCode, "The specified resource was not found", "")
	case statusCode == http.StatusConflict:
		return NewError(statusCode, "Parameter is invalid or incorrect.", "")
	default:
		return nil
	}
}

func GetUrl(address string, apiKey string, resource string, query proto.Message) string {
	sb := strings.Builder{}
	sb.WriteString(address)
	sb.WriteString(resource)
	sb.WriteString("?apiKey=")
	sb.WriteString(apiKey)
	if query != nil {
		m := proto.MessageReflect(query)
		md := m.Descriptor()
		fds := md.Fields()
		for i := 0; i < fds.Len(); {
			fd := fds.Get(i)
			i++
			v := m.Get(fd)
			if m.Has(fd) {
				sb.WriteString(fmt.Sprintf("#%v=%v", fd.JSONName(), v.Interface()))
			}
		}
	}
	return sb.String()
}

type Error struct {
	ErrorCode    int
	ErrorMessage string
	RequestId    string
}

func NewError(errorCode int, errorMessage string, requestId string) *Error {
	return &Error{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		RequestId:    requestId,
	}
}

func (self *Error) Error() string {
	return fmt.Sprintf("error: %v, requestId: %v, errorCode: %v, errorDescription: %v", self.ErrorMessage, self.RequestId, self.ErrorCode, http.StatusText(self.ErrorCode))
}
