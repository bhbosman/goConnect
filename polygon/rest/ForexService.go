package rest

import (
	"context"
	"github.com/bhbosman/goConnect/polygon/rest/internal"
	"github.com/bhbosman/goMessages/polygon/rest"
	"net/http"
)

type ForexService struct {
	baseService
}

func (self *ForexService) GetForexPreviousClose(ctx context.Context, query *rest.GetForexPreviousCloseQueryRequest) (*rest.GetForexPreviousCloseResponse, error) {
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

func (self *ForexService) GetForexAggregatesBars(ctx context.Context, query *rest.GetForexAggregatesBarsQueryRequest) (*rest.GetForexAggregatesBarsResponse, error) {
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

func (self *ForexService) GetForexGroupedDailyBars(ctx context.Context, query *rest.GetForexGroupedDailyBarsQueryRequest) (*rest.GetForexGroupedDailyBarsResponse, error) {
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

func (self *ForexService) GetForexHistoricForexTicks(ctx context.Context, query *rest.GetForexHistoricForexTicksQueryRequest) (*rest.GetForexHistoricForexTicksResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/ticks/stocks/trades/{ticker}/{date}", nil),
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

func (self *ForexService) GetForexRealTimeCurrencyConversion(ctx context.Context, query *rest.GetForexRealTimeCurrencyConversionQueryRequest) (*rest.GetForexRealTimeCurrencyConversionResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/conversion/{from}/{to}", nil),
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

func (self *ForexService) GetForexLastQuoteForCurrencyPair(ctx context.Context, query *rest.GetForexLastQuoteForCurrencyPairQueryRequest) (*rest.GetForexLastQuoteForCurrencyPairResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/last_quote/currencies/{from}/{to}", nil),
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

func (self *ForexService) GetForexSnapshotAllTickers(ctx context.Context, query *rest.GetForexSnapshotAllTickersQueryRequest) (*rest.GetForexSnapshotAllTickersResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/global/markets/forex/tickers", nil),
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

func (self *ForexService) GetForexSnapshotGainersLosers(ctx context.Context, query *rest.GetForexSnapshotGainersLosersQueryRequest) (*rest.GetForexSnapshotGainersLosersResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/global/markets/forex/{direction}", nil),
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
