package rest

import (
	"context"
	"github.com/bhbosman/goConnect/polygon/rest/internal"
	"github.com/bhbosman/goMessages/polygon/rest"
	"net/http"
)

type StockService struct {
	baseService
}

func (self *StockService) GetEquityExchanges(ctx context.Context, query *rest.GetEquityExchangesQueryRequest) (*rest.GetEquityExchangesResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/meta/exchanges", nil),
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

func (self *StockService) GetEquityPreviousClose(ctx context.Context, query *rest.GetEquityPreviousCloseQueryRequest) (*rest.GetEquityPreviousCloseResponse, error) {
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

func (self *StockService) GetEquityAggregatesBars(ctx context.Context, query *rest.GetEquityAggregatesBarsQueryRequest) (*rest.GetEquityAggregatesBarsResponse, error) {
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

func (self *StockService) GetEquityGroupedDailyBars(ctx context.Context, query *rest.GetEquityGroupedDailyBarsQueryRequest) (*rest.GetEquityGroupedDailyBarsResponse, error) {
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

func (self *StockService) GetEquityHistoricTrades(ctx context.Context, query *rest.GetEquityHistoricTradesQueryRequest) (*rest.GetEquityHistoricTradesResponse, error) {
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

func (self *StockService) GetEquityHistoricQuotesNBBO(ctx context.Context, query *rest.GetEquityHistoricQuotesNBBOQueryRequest) (*rest.GetEquityHistoricQuotesNBBOResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/ticks/stocks/nbbo/{ticker}/{date}", nil),
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

func (self *StockService) GetEquityLastTradeForSymbol(ctx context.Context, query *rest.GetEquityLastTradeForSymbolQueryRequest) (*rest.GetEquityLastTradeForSymbolResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/last/stocks/{symbol}", nil),
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

func (self *StockService) GetEquityLastQuoteForSymbol(ctx context.Context, query *rest.GetEquityLastQuoteForSymbolQueryRequest) (*rest.GetEquityLastQuoteForSymbolResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/last_quote/stocks/{symbol}", nil),
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

func (self *StockService) GetEquityDailyOpenClose(ctx context.Context, query *rest.GetEquityDailyOpenCloseQueryRequest) (*rest.GetEquityDailyOpenCloseResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/open-close/{symbol}/{date}", nil),
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

func (self *StockService) GetEquityConditionMappings(ctx context.Context, query *rest.GetEquityConditionMappingsQueryRequest) (*rest.GetEquityConditionMappingsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/meta/conditions/{ticktype}", nil),
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

func (self *StockService) GetEquitySnapshotAllTickers(ctx context.Context, query *rest.GetEquitySnapshotAllTickersQueryRequest) (*rest.GetEquitySnapshotAllTickersResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/us/markets/stocks/tickers", nil),
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

func (self *StockService) GetEquitySnapshotSingleTicker(ctx context.Context, query *rest.GetEquitySnapshotSingleTickerQueryRequest) (*rest.GetEquitySnapshotSingleTickerResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/us/markets/stocks/tickers/{ticker}", nil),
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

func (self *StockService) GetEquitySnapshotGainersLosers(ctx context.Context, query *rest.GetEquitySnapshotGainersLosersQueryRequest) (*rest.GetEquitySnapshotGainersLosersResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/snapshot/locale/us/markets/stocks/{direction}", nil),
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
