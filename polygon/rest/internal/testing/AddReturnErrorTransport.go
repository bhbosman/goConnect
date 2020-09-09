package testing

import "net/http"

type AddReturnErrorTransport struct {
	Transport http.RoundTripper
}

func (t *AddReturnErrorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	newReq := *req
	args := newReq.URL.Query()
	args.Set("RETURN_ERROR", "RETURN_ERROR")
	newReq.URL.RawQuery = args.Encode()
	return rt.RoundTrip(&newReq)
}
