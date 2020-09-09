package rest_test

import (
	"context"
	"github.com/bhbosman/goConnect/polygon/rest"
	testing2 "github.com/bhbosman/goConnect/polygon/rest/internal/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolygonReferenceService(t *testing.T) {
	httpServer := testing2.PolygonTestServerForReferenceService()
	defer func() {
		_ = httpServer.Close()
	}()
	t.Run("Get Tickers", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetTickers(context.Background(), nil)
			assert.NoError(t, err)
			if !assert.Len(t, response.Symbols, 2) {
				return
			}
			if !assert.Equal(t, "AAPL", response.Symbols[0].Ticker) {
				return
			}
			if !assert.Equal(t, "$AEDAUD", response.Symbols[1].Ticker) {
				return
			}
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetTickers(context.Background(), nil)
			assert.Error(t, err)
		})
	})

	t.Run("Get Ticker Types", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			result, err := restService.GetTickerTypes(context.Background())
			assert.NoError(t, err)
			if !assert.Equal(t, 3, len(result.IndexTypes)) {
				return
			}
			if !assert.Equal(t, 4, len(result.Types)) {
				return
			}
		})
		t.Run("With error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetTickerTypes(context.Background())
			assert.Error(t, err)
		})
	})

	t.Run("Get Ticker Details", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetTickerDetails(context.Background(), "AAPL")
			if !assert.NoError(t, err) {
				return
			}
			if !assert.Equal(t, "AAPL", response.Symbol) {
				return
			}
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetTickerDetails(context.Background(), "Anything")
			assert.Error(t, err)
		})
	})

	t.Run("Get Tickers News", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetTickerNews(context.Background(), "AAPL", nil)
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetTickerNews(context.Background(), "Anything", nil)
			assert.Error(t, err)
		})
	})

	t.Run("Get Markets", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetMarkets(context.Background())
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetMarkets(context.Background())
			assert.Error(t, err)
		})
	})

	t.Run("Get Locale", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetLocales(context.Background())
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetLocales(context.Background())
			assert.Error(t, err)
		})
	})

	t.Run("Get stock splits", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetStockSplits(context.Background(), "AA")
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetStockSplits(context.Background(), "2222")
			assert.Error(t, err)
		})
	})

	t.Run("Get stock Dividend", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetStockDividends(context.Background(), "AA")
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetStockDividends(context.Background(), "2222")
			assert.Error(t, err)
		})
	})

	t.Run("Get stock Financials", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetStockFinancials(context.Background(), "AA", nil)
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetStockFinancials(context.Background(), "2222", nil)
			assert.Error(t, err)
		})
	})

	t.Run("Get market status", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetMarketStatus(context.Background())
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetMarketStatus(context.Background())
			assert.Error(t, err)
		})
	})

	t.Run("Get market upcoming", func(t *testing.T) {
		t.Run("With no error from server", func(t *testing.T) {
			restService := rest.NewPolygonRestService(testing2.TestAddress, "1234")
			response, err := restService.GetMarketHolidays(context.Background())
			if !assert.NoError(t, err) {
				return
			}
			assert.NotNil(t, response)
		})
		t.Run("With server error", func(t *testing.T) {
			restService := rest.NewPolygonRestServiceWithTransport(testing2.TestAddress, "1234", &testing2.AddReturnErrorTransport{})
			_, err := restService.GetMarketHolidays(context.Background())
			assert.Error(t, err)
		})
	})
}
