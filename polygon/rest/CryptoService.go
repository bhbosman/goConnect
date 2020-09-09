package rest

import (
	"context"
	"github.com/bhbosman/goConnect/polygon/rest/internal"
	"github.com/bhbosman/goMessages/polygon/rest"
	"net/http"
)

type CryptoService struct {
	baseService
}

func (self *ReferenceService) GetCryptoPreviousClose(ctx context.Context, query *rest.GetCryptoPreviousCloseQueryRequest) (*rest.GetCryptoPreviousCloseResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/aggs/ticker/{ticker}/prev", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoAggregatesBars(ctx context.Context, query *rest.GetCryptoAggregatesBarsQueryRequest) (*rest.GetCryptoAggregatesBarsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/aggs/ticker/{ticker}/range/{multiplier}/{timespan}/{from}/{to}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoGroupedDailyBars(ctx context.Context, query *rest.GetCryptoGroupedDailyBarsQueryRequest) (*rest.GetCryptoGroupedDailyBarsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/aggs/grouped/locale/{locale}/market/{market}/{date}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoCryptoExchanges(ctx context.Context, query *rest.GetCryptoCryptoExchangesQueryRequest) (*rest.GetCryptoCryptoExchangesResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/meta/crypto-exchanges", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoLastTradeForCryptoPair(ctx context.Context, query *rest.GetCryptoLastTradeForCryptoPairQueryRequest) (*rest.GetCryptoLastTradeForCryptoPairResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/last/crypto/{from}/{to}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoDailyOpenClose(ctx context.Context, query *rest.GetCryptoDailyOpenCloseQueryRequest) (*rest.GetCryptoDailyOpenCloseResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/open-close/crypto/{from}/{to}/{date}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoHistoricCryptoTrades(ctx context.Context, query *rest.GetCryptoHistoricCryptoTradesQueryRequest) (*rest.GetCryptoHistoricCryptoTradesResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/historic/crypto/{from}/{to}/{date}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoSnapshotAllTickers(ctx context.Context, query *rest.GetCryptoSnapshotAllTickersQueryRequest) (*rest.GetCryptoSnapshotAllTickersResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/global/markets/crypto/tickers", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoSnapshotSingleTicker(ctx context.Context, query *rest.GetCryptoSnapshotSingleTickerQueryRequest) (*rest.GetCryptoSnapshotSingleTickerResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/global/markets/crypto/tickers/{ticker}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoSnapshotSingleTickerFullBookL2(ctx context.Context, query *rest.GetCryptoSnapshotSingleTickerFullBookL2QueryRequest) (*rest.GetCryptoSnapshotSingleTickerFullBookL2Response, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/global/markets/crypto/tickers/{ticker}/book", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (self *ReferenceService) GetCryptoSnapshotGainersLosers(ctx context.Context, query *rest.GetCryptoSnapshotGainersLosersQueryRequest) (*rest.GetCryptoSnapshotGainersLosersResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/global/markets/crypto/{direction}", nil),
		nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	err = internal.ReadCommonError(response.StatusCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
