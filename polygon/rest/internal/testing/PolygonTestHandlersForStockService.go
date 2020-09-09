package testing

import (
	"github.com/gorilla/mux"
	"net/http"
	url2 "net/url"
)

func PolygonTestHandlersForStockService() http.Handler {
	result := mux.NewRouter()
	result.HandleFunc("/v1/meta/exchanges", GetExchanges)
	result.HandleFunc("/v2/aggs/ticker/{ticker}/prev", GetPreviousClose)
	result.HandleFunc("/v2/aggs/ticker/{ticker}/range/{multiplier}/{timespan}/{from}/{to}", GetAggregatesBars)
	result.HandleFunc("/v2/aggs/grouped/locale/{locale}/market/{market}/{date}", GetGroupedDailyBars)
	result.HandleFunc("/v2/ticks/stocks/trades/{ticker}/{date}", GetHistoricTrades)
	result.HandleFunc("/v2/ticks/stocks/trades/{ticker}/{date}", GetHistoricQuotesNBBO)
	result.HandleFunc("/v1/last/stocks/{symbol}", GetLastTradeForSymbol)
	result.HandleFunc("/v1/last_quote/stocks/{symbol}", GetLastQuoteForSymbol)
	result.HandleFunc("/v1/open-close/{symbol}/{date}", GetDailyOpenClose)
	result.HandleFunc("/v1/meta/conditions/{ticktype}", GetConditionMappings)
	result.HandleFunc("/v2/snapshot/locale/us/markets/stocks/tickers", GetSnapshotAllTickers)
	result.HandleFunc("/v2/snapshot/locale/us/markets/stocks/tickers/{ticker}", GetSnapshotSingleTicker)
	result.HandleFunc("/v2/snapshot/locale/us/markets/stocks/{direction}", GetSnapshotGainersLosers)

	return result
}

func PolygonTestServerForStockService() *http.Server {
	url, _ := url2.Parse(TestAddress)
	result := &http.Server{
		Addr:    url.Host,
		Handler: PolygonTestHandlersForStockService(),
	}
	go func() {
		_ = result.ListenAndServe()
	}()
	return result

}

func GetSnapshotGainersLosers(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		
		{
		  "status": "OK",
		  "tickers": [
			{
			  "ticker": "AAPL",
			  "day": {
				"c": 0.2907,
				"h": 0.2947,
				"l": 0.2901,
				"o": 0.2905,
				"v": 1432
			  },
			  "lastTrade": {
				"c1": 14,
				"c2": 12,
				"c3": 0,
				"c4": 0,
				"e": 12,
				"p": 172.17,
				"s": 50,
				"t": 1517529601006
			  },
			  "lastQuote": {
				"p": 120,
				"s": 5,
				"P": 121,
				"S": 3,
				"t": 1547787608999000000
			  },
			  "min": {
				"c": 0.2907,
				"h": 0.2947,
				"l": 0.2901,
				"o": 0.2905,
				"v": 1432
			  },
			  "prevDay": {
				"c": 0.2907,
				"h": 0.2947,
				"l": 0.2901,
				"o": 0.2905,
				"v": 1432
			  },
			  "todaysChange": 0.001,
			  "todaysChangePerc": 2.55,
			  "updated": 1547787608999
			}
		  ]
		}


	`))
}

func GetSnapshotSingleTicker(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "status": "OK",
		  "ticker": {
			"ticker": "AAPL",
			"day": {
			  "c": 0.2907,
			  "h": 0.2947,
			  "l": 0.2901,
			  "o": 0.2905,
			  "v": 1432
			},
			"lastTrade": {
			  "c1": 14,
			  "c2": 12,
			  "c3": 0,
			  "c4": 0,
			  "e": 12,
			  "p": 172.17,
			  "s": 50,
			  "t": 1517529601006
			},
			"lastQuote": {
			  "p": 120,
			  "s": 5,
			  "P": 121,
			  "S": 3,
			  "t": 1547787608999000000
			},
			"min": {
			  "c": 0.2907,
			  "h": 0.2947,
			  "l": 0.2901,
			  "o": 0.2905,
			  "v": 1432
			},
			"prevDay": {
			  "c": 0.2907,
			  "h": 0.2947,
			  "l": 0.2901,
			  "o": 0.2905,
			  "v": 1432
			},
			"todaysChange": 0.001,
			"todaysChangePerc": 2.55,
			"updated": 1547787608999
		  }
		}
`))
}

func GetSnapshotAllTickers(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
				{
				  "status": "OK",
				  "tickers": [
					{
					  "ticker": "AAPL",
					  "day": {
						"c": 0.2907,
						"h": 0.2947,
						"l": 0.2901,
						"o": 0.2905,
						"v": 1432
					  },
					  "lastTrade": {
						"c1": 14,
						"c2": 12,
						"c3": 0,
						"c4": 0,
						"e": 12,
						"p": 172.17,
						"s": 50,
						"t": 1517529601006
					  },
					  "lastQuote": {
						"p": 120,
						"s": 5,
						"P": 121,
						"S": 3,
						"t": 1547787608999000000
					  },
					  "min": {
						"c": 0.2907,
						"h": 0.2947,
						"l": 0.2901,
						"o": 0.2905,
						"v": 1432
					  },
					  "prevDay": {
						"c": 0.2907,
						"h": 0.2947,
						"l": 0.2901,
						"o": 0.2905,
						"v": 1432
					  },
					  "todaysChange": 0.001,
					  "todaysChangePerc": 2.55,
					  "updated": 1547787608999
					}
				  ]
				}
			`))
}

func GetConditionMappings(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		
		{
		  "1": "Regular",
		  "2": "Acquisition",
		  "3": "AveragePrice",
		  "4": "AutomaticExecution"
		}

`))
}

func GetDailyOpenClose(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "status": "OK",
		  "from": "2020-06-03",
		  "symbol": "NVDA",
		  "open": 352.89,
		  "high": 352.89,
		  "low": 352.89,
		  "close": 352.89,
		  "volume": 9194990,
		  "preMarket": 355,
		  "afterHours": 350.24
		}
	`))
}

func GetLastQuoteForSymbol(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "status": "success",
		  "symbol": "AAPL",
		  "last": {
			"askprice": 159.59,
			"asksize": 2,
			"askexchange": 11,
			"bidprice": 159.45,
			"bidsize": 20,
			"bidexchange": 12,
			"timestamp": 1518086601843
		  }
		}
	`))
}

func GetLastTradeForSymbol(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "status": "success",
		  "symbol": "AAPL",
		  "last": {
			"price": 159.59,
			"size": 20,
			"exchange": 11,
			"cond1": 14,
			"cond2": 16,
			"cond3": 0,
			"cond4": 0,
			"timestamp": 1518086464720
		  }
		}
`))
}

func GetHistoricQuotesNBBO(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
			{
			  "results_count": 10,
			  "db_latency": 2,
			  "success": true,
			  "ticker": "AAPL",
			  "results": [
				{
				  "T": "AAPL",
				  "t": 1547787608999125800,
				  "y": 1547787608999125800,
				  "f": 1547787608999125800,
				  "q": 23547,
				  "c": [
					{}
				  ],
				  "i": [
					{}
				  ],
				  "p": 223.001,
				  "x": 11,
				  "s": 100,
				  "P": 223.001,
				  "X": 11,
				  "S": 100,
				  "z": 1
				}
			  ]
			}
		`))
}

func GetHistoricTrades(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "results_count": 10,
		  "db_latency": 2,
		  "success": true,
		  "ticker": "AAPL",
		  "results": [
			{
			  "T": "AAPL",
			  "t": 1547787608999125800,
			  "y": 1547787608999125800,
			  "f": 1547787608999125800,
			  "q": 23547,
			  "i": "00MGON",
			  "x": 11,
			  "s": 100,
			  "c": [
				{}
			  ],
			  "p": 223.001,
			  "z": 1
			}
		  ]
		}
`))
}

func GetGroupedDailyBars(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "ticker": "AAPL",
		  "status": "OK",
		  "adjusted": true,
		  "queryCount": 55,
		  "resultsCount": 2,
		  "results": [
			{
			  "T": "AAPL",
			  "v": 31315282,
			  "o": 102.87,
			  "c": 103.74,
			  "h": 103.82,
			  "l": 102.65,
			  "t": 1549314000000,
			  "n": 4
			}
		  ]
		}`))
}

func GetAggregatesBars(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		{
		  "ticker": "AAPL",
		  "status": "OK",
		  "adjusted": true,
		  "queryCount": 55,
		  "resultsCount": 2,
		  "results": [
			{
			  "T": "AAPL",
			  "v": 31315282,
			  "o": 102.87,
			  "c": 103.74,
			  "h": 103.82,
			  "l": 102.65,
			  "t": 1549314000000,
			  "n": 4
			}
		  ]
		}`))
}

func GetExchanges(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte(`
		[
		  {
			"id": 1,
			"type": "exchange",
			"market": "equities",
			"mic": "XASE",
			"name": "NYSE American (AMEX)",
			"tape": "A"
		  },
		  {
			"id": 2,
			"type": "exchange",
			"market": "equities",
			"mic": "XBOS",
			"name": "NASDAQ OMX BX",
			"tape": "B"
		  },
		  {
			"id": 15,
			"type": "exchange",
			"market": "equities",
			"mic": "IEXG",
			"name": "IEX",
			"tape": "V"
		  },
		  {
			"id": 16,
			"type": "TRF",
			"market": "equities",
			"mic": "XCBO",
			"name": "Chicago Board Options Exchange",
			"tape": "W"
		  }
		]`))

}

func GetPreviousClose(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte(`

		{
		  "ticker": "AAPL",
		  "status": "OK",
		  "adjusted": true,
		  "queryCount": 55,
		  "resultsCount": 2,
		  "results": [
			{
			  "T": "AAPL",
			  "v": 31315282,
			  "o": 102.87,
			  "c": 103.74,
			  "h": 103.82,
			  "l": 102.65,
			  "t": 1549314000000,
			  "n": 4
			}
		  ]
		}

`))
}
