package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bhbosman/goConnect/polygon/rest/internal"
	"github.com/bhbosman/goMessages/polygon/rest"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
)

type ReferenceService struct {
	baseService
}

func (self *ReferenceService) GetTickers(ctx context.Context, query *rest.GetTickerQueryRequest) (*rest.GetTickerResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/reference/tickers", query),
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
	switch {
	case response.StatusCode == http.StatusOK:
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		rawDataForItems, err := internal.ReadJsonArray(response.Body)
		if err != nil {
			return nil, err
		}
		tickerResponse := &rest.GetTickerResponse{}
		tickerResponse.Symbols = make([]*rest.Symbol, len(rawDataForItems))
		for index, rawDataForItem := range rawDataForItems {
			symbol := &rest.Symbol{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), symbol)
			if err != nil {
				return nil, err
			}
			tickerResponse.Symbols[index] = symbol
		}
		return tickerResponse, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetTickerTypes(ctx context.Context) (*rest.GetTickerTypesResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/reference/types", nil),
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

	switch {
	case response.StatusCode == http.StatusOK:
		data := struct {
			Status    string          `json:"status"`
			RequestId string          `json:"request_id"`
			Error     string          `json:"error"`
			Results   json.RawMessage `json:"results"`
		}{}
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		result := &rest.GetTickerTypesResponse{}
		err = unmarshaler.Unmarshal(bytes.NewReader(data.Results), result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetTickerDetails(ctx context.Context, symbol string) (*rest.GetTickerDetailsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, fmt.Sprintf("/v1/meta/symbols/%v/company", symbol), nil),
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

	switch {
	case response.StatusCode == http.StatusOK:
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		result := &rest.GetTickerDetailsResponse{}
		err = unmarshaler.Unmarshal(response.Body, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetTickerNews(ctx context.Context, symbol string, query *rest.GetTickerNewsQueryRequest) (*rest.GetTickerNewsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, fmt.Sprintf("/v1/meta/symbols/%v/news", symbol), query),
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
	switch {
	case response.StatusCode == http.StatusOK:
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		rawDataForItems, err := internal.ReadJsonArray(response.Body)
		if err != nil {
			return nil, err
		}
		result := &rest.GetTickerNewsResponse{}
		result.News = make([]*rest.News, len(rawDataForItems))
		for index, rawDataForItem := range rawDataForItems {
			news := &rest.News{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), news)
			if err != nil {
				return nil, err
			}
			result.News[index] = news
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetMarkets(ctx context.Context) (*rest.GetMarketsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/reference/markets", nil),
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

	switch {
	case response.StatusCode == http.StatusOK:
		data := struct {
			Status       string          `json:"status"`
			ResultsArray json.RawMessage `json:"results"`
		}{}
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
		arrayData, err := internal.ReadJsonArray(bytes.NewBuffer(data.ResultsArray))
		if err != nil {
			return nil, err
		}
		result := &rest.GetMarketsResponse{}
		result.Markets = make([]*rest.Market, len(arrayData))
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		for index, rawDataForItem := range arrayData {
			market := &rest.Market{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), market)
			if err != nil {
				return nil, err
			}
			result.Markets[index] = market
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetLocales(ctx context.Context) (*rest.GetLocalesResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v2/reference/locales", nil),
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
	switch {
	case response.StatusCode == http.StatusOK:
		data := struct {
			Status       string          `json:"status"`
			ResultsArray json.RawMessage `json:"results"`
		}{}
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
		arrayData, err := internal.ReadJsonArray(bytes.NewBuffer(data.ResultsArray))
		if err != nil {
			return nil, err
		}
		result := &rest.GetLocalesResponse{}
		result.Locales = make([]*rest.Locale, len(arrayData))
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		for index, rawDataForItem := range arrayData {
			locale := &rest.Locale{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), locale)
			if err != nil {
				return nil, err
			}
			result.Locales[index] = locale
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}

}

func (self *ReferenceService) GetStockSplits(ctx context.Context, symbol string) (*rest.GetStockSplitsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, fmt.Sprintf("/v2/reference/splits/%v", symbol), nil),
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
	switch {
	case response.StatusCode == http.StatusOK:
		data := struct {
			Status       string          `json:"status"`
			ResultsArray json.RawMessage `json:"results"`
		}{}
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
		arrayData, err := internal.ReadJsonArray(bytes.NewBuffer(data.ResultsArray))
		if err != nil {
			return nil, err
		}
		result := &rest.GetStockSplitsResponse{}
		result.Splits = make([]*rest.Split, len(arrayData))
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		for index, rawDataForItem := range arrayData {
			split := &rest.Split{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), split)
			if err != nil {
				return nil, err
			}
			result.Splits[index] = split
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetStockDividends(ctx context.Context, symbol string) (*rest.GetStockDividendsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, fmt.Sprintf("/v2/reference/dividends/%v", symbol), nil),
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
	switch {
	case response.StatusCode == http.StatusOK:
		data := struct {
			Status       string          `json:"status"`
			ResultsArray json.RawMessage `json:"results"`
		}{}
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
		arrayData, err := internal.ReadJsonArray(bytes.NewBuffer(data.ResultsArray))
		if err != nil {
			return nil, err
		}
		result := &rest.GetStockDividendsResponse{}
		result.Dividends = make([]*rest.Dividend, len(arrayData))
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		for index, rawDataForItem := range arrayData {
			dividend := &rest.Dividend{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), dividend)
			if err != nil {
				return nil, err
			}
			result.Dividends[index] = dividend
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetStockFinancials(ctx context.Context, symbol string, query *rest.GetStockFinancialsQueryRequest) (*rest.GetStockFinancialsResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(
			self.address,
			self.apiKey,
			fmt.Sprintf("/v2/reference/financials/%v", symbol),
			query),
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
	switch {
	case response.StatusCode == http.StatusOK:
		data := struct {
			Status       string          `json:"status"`
			ResultsArray json.RawMessage `json:"results"`
		}{}
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
		arrayData, err := internal.ReadJsonArray(bytes.NewBuffer(data.ResultsArray))
		if err != nil {
			return nil, err
		}
		result := &rest.GetStockFinancialsResponse{}
		result.Financials = make([]*rest.Financial, len(arrayData))
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		for index, rawDataForItem := range arrayData {
			financial := &rest.Financial{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), financial)
			if err != nil {
				return nil, err
			}
			result.Financials[index] = financial
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetMarketStatus(ctx context.Context) (*rest.GetMarketStatusResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/marketstatus/now", nil),
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
	switch {
	case response.StatusCode == http.StatusOK:
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		result := &rest.GetMarketStatusResponse{}
		err = unmarshaler.Unmarshal(response.Body, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}
}

func (self *ReferenceService) GetMarketHolidays(ctx context.Context) (*rest.GetMarketHolidaysResponse, error) {
	var req, err = http.NewRequestWithContext(
		ctx,
		"GET",
		internal.GetUrl(self.address, self.apiKey, "/v1/marketstatus/upcoming", nil),
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
	switch {
	case response.StatusCode == http.StatusOK:
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
			AnyResolver:        nil,
		}
		rawDataForItems, err := internal.ReadJsonArray(response.Body)
		if err != nil {
			return nil, err
		}
		result := &rest.GetMarketHolidaysResponse{}
		result.MarketHolidays = make([]*rest.MarketHoliday, len(rawDataForItems))
		for index, rawDataForItem := range rawDataForItems {
			marketHoliday := &rest.MarketHoliday{}
			err = unmarshaler.Unmarshal(bytes.NewReader(rawDataForItem), marketHoliday)
			if err != nil {
				return nil, err
			}
			result.MarketHolidays[index] = marketHoliday
		}
		return result, nil
	default:
		return nil, internal.ReadJsonError(response.Body, response.StatusCode)
	}

}

func NewPolygonRestService(address string, apiKey string) *ReferenceService {
	return &ReferenceService{
		baseService: baseService{
			apiKey: apiKey,
			client: http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       0,
			},
			address: address,
		},
	}
}

func NewPolygonRestServiceWithTransport(address string, apiKey string, transport http.RoundTripper) *ReferenceService {
	return &ReferenceService{
		baseService: baseService{
			apiKey: apiKey,
			client: http.Client{
				Transport:     transport,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       0,
			},
			address: address,
		},
	}
}
