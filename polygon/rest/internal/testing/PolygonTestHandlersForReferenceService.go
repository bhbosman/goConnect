package testing

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	url2 "net/url"
)

func PolygonTestHandlersForReferenceService() http.Handler {
	result := mux.NewRouter()
	result.HandleFunc("/v2/reference/tickers", GetTickerHandler)
	result.HandleFunc("/v2/reference/types", GetTickerTypes)
	result.HandleFunc("/v1/meta/symbols/{symbol}/company", GetTickerDetail)
	result.HandleFunc("/v1/meta/symbols/{symbol}/news", GetTickerNews)
	result.HandleFunc("/v2/reference/markets", GetMarkets)
	result.HandleFunc("/v2/reference/locales", GetLocale)
	result.HandleFunc("/v2/reference/splits/{symbol}", GetStockSplits)
	result.HandleFunc("/v2/reference/dividends/{symbol}", GetStockDividends)
	result.HandleFunc("/v2/reference/financials/{symbol}", GetStockFinancials)
	result.HandleFunc("/v1/marketstatus/now", GetMarketStatus)
	result.HandleFunc("/v1/marketstatus/upcoming", GetMarketStatusUpcoming)
	return result
}

const TestAddress = "http://127.0.0.1:9090"

func PolygonTestServerForReferenceService() *http.Server {
	url, _ := url2.Parse(TestAddress)
	result := &http.Server{
		Addr:    url.Host,
		Handler: PolygonTestHandlersForReferenceService(),
	}
	go func() {
		_ = result.ListenAndServe()
	}()
	return result

}

func GetTickerHandler(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`[
			 {
			   "ticker": "AAPL",
			   "name": "Apple Inc.",
			   "market": "STOCKS",
			   "locale": "US",
			   "currency": "USD",
			   "active": true,
			   "primaryExch": "NGS",
			   "type": "cs",
			   "codes": {
				 "cik": "0000320193",
				 "figiuid": "EQ0010169500001000",
				 "scfigi": "BBG001S5N8V8",
				 "cfigi": "BBG000B9XRY4",
				 "figi": "BBG000B9Y5X2"
			   },
			   "updated": "2019-01-15T05:21:28.437Z",
			   "url": "https://api.polygon.io/v2/reference/tickers/AAPL"
			 },
			 {
			   "ticker": "$AEDAUD",
			   "name": "United Arab Emirates dirham - Australian dollar",
			   "market": "FX",
			   "locale": "G",
			   "currency": "AUD",
			   "active": true,
			   "primaryExch": "FX",
			   "updated": "2019-01-25T00:00:00.000Z",
			   "attrs": {
				 "currencyName": "Australian dollar,",
				 "currency": "AUD,",
				 "baseName": "United Arab Emirates dirham,",
				 "base": "AED"
			   },
			   "url": "https://api.polygon.io/v2/tickers/$AEDAUD"
			 }
			]`))
	}
}

func GetTickerTypes(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(`
			{
			  "status": "OK",
			  "results": {
				"types": {
				  "CS": "Common Stock",
				  "ADR": "American Depository Receipt",
				  "NVDR": "Non-Voting Depository Receipt",
				  "GDR": "Global Depositary Receipt"
				},
				"indexTypes": {
				  "INDEX": "Index",
				  "ETF": "Exchange Traded Fund (ETF)",
				  "ETN": "Exchange Traded Note (ETN)"
				}
			  }
			}`))
	}
}

func GetTickerDetail(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
					{
					  "logo": "https://s3.polygon.io/logos/aapl/logo.png",
					  "exchange": "Nasdaq Global Select",
					  "name": "Apple Inc.",
					  "symbol": "AAPL",
					  "cik": "0000320193",
					  "bloomberg": "EQ0010169500001000",
					  "lei": "HWUPKR0MPOU8FGXBT394",
					  "sic": 3571,
					  "country": "us",
					  "industry": "Computer Hardware",
					  "sector": "Technology",
					  "marketcap": 815604985500,
					  "employees": 116000,
					  "phone": "(408) 996-1010",
					  "ceo": "Tim Cook",
					  "url": "http://www.apple.com",
					  "description": "Apple Inc. designs, manufactures, and markets mobile communication and media devices, personal computers, and portable digital music players to consumers...\n",
					  "similar": [
						"MSFT",
						"IBM",
						"GOOGL"
					  ],
					  "tags": [
						"Technology",
						"Consumer Electronics",
						"Computer Hardware"
					  ]
					}

				`))
	}
}

func GetTickerNews(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
					[
					  {
						"symbols": [
						  "MSFT",
						  "AAPL",
						  "IBM"
						],
						"title": "Goldman in talks to finance iPhones - WSJ",
						"url": "https://seekingalpha.com/news/3328826-goldman-talks-finance-iphones-wsj",
						"source": "SeekingAlpha",
						"summary": "Continuing its push into more traditional areas of consumer finance, Goldman Sachs (NYSE:GS) is reportedly in talks with Apple (NASDAQ:AAPL) to finance iPhone purchases.Buyers theoretically would be a...",
						"image": "https://static.seekingalpha.com/assets/og_image_410-b8960ce31ec84f7f12dba11a09fc1849b69b234e0f5f39d7c62f46f8692e58a5.png",
						"timestamp": "2018-02-07T12:48:47.000Z",
						"keywords": [
						  "financial services",
						  "aapl",
						  "investing",
						  "bsiness news",
						  "mobile"
						]
					  }
					]
				`))
	}
}

func GetMarkets(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
				{
				  "status": "OK",
				  "results": [
					{
					  "market": "STOCKS",
					  "desc": "Stocks / Equities / ETFs"
					},
					{
					  "market": "INDICES",
					  "desc": "Indices"
					},
					{
					  "market": "MF",
					  "desc": "Mutual Funds"
					}
				  ]
				}
			`))
	}
}

func GetLocale(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
					{
					  "status": "OK",
					  "results": [
						{
						  "locale": "US",
						  "name": "United States"
						},
						{
						  "locale": "GB",
						  "name": "Great Britain / UK"
						},
						{
						  "locale": "CA",
						  "name": "Canada"
						}
					  ]
					}
			`))
	}
}

func GetStockSplits(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
					{
					  "status": "OK",
					  "count": 1,
					  "results": [
						{
						  "ticker": "AAPL",
						  "exDate": "1999-03-28",
						  "paymentDate": "1999-03-28",
						  "recordDate": "1999-03-28",
						  "declaredDate": "1999-03-28",
						  "ratio": 0.142857,
						  "tofactor": 7,
						  "forfactor": 1
						}
					  ]
					}
			`))
	}
}

func GetStockDividends(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
					{
					  "status": "OK",
					  "count": 1,
					  "results": [
						{
						  "symbol": "AAPL",
						  "type": "Dividend income",
						  "exDate": "2016-11-03T04:00:00.000Z",
						  "paymentDate": "2016-11-03T04:00:00.000Z",
						  "recordDate": "2016-11-03T04:00:00.000Z",
						  "declaredDate": "2016-11-03T04:00:00.000Z",
						  "amount": 0.57,
						  "qualified": "Q",
						  "flag": "YE"
						}
					  ]
					}
			`))
	}
}

func GetStockFinancials(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(
			`
				{
				  "status": "OK",
				  "count": 1,
				  "results": [
					{
					  "ticker": "AAPL",
					  "period": "Q",
					  "calendarDate": "2019-03-31",
					  "reportPeriod": "2019-03-31",
					  "updated": "1999-03-28",
					  "accumulatedOtherComprehensiveIncome": 0,
					  "assets": 0,
					  "assetsAverage": 0,
					  "assetsCurrent": 0,
					  "assetTurnover": 0,
					  "assetsNonCurrent": 0,
					  "bookValuePerShare": 0,
					  "capitalExpenditure": 0,
					  "cashAndEquivalents": 0,
					  "cashAndEquivalentsUSD": 0,
					  "costOfRevenue": 0,
					  "consolidatedIncome": 0,
					  "currentRatio": 0,
					  "debtToEquityRatio": 0,
					  "debt": 0,
					  "debtCurrent": 0,
					  "debtNonCurrent": 0,
					  "debtUSD": 0,
					  "deferredRevenue": 0,
					  "depreciationAmortizationAndAccretion": 0,
					  "deposits": 0,
					  "dividendYield": 0,
					  "dividendsPerBasicCommonShare": 0,
					  "earningBeforeInterestTaxes": 0,
					  "earningsBeforeInterestTaxesDepreciationAmortization": 0,
					  "EBITDAMargin": 0,
					  "earningsBeforeInterestTaxesDepreciationAmortizationUSD": 0,
					  "earningBeforeInterestTaxesUSD": 0,
					  "earningsBeforeTax": 0,
					  "earningsPerBasicShare": 0,
					  "earningsPerDilutedShare": 0,
					  "earningsPerBasicShareUSD": 0,
					  "shareholdersEquity": 0,
					  "averageEquity": 0,
					  "shareholdersEquityUSD": 0,
					  "enterpriseValue": 0,
					  "enterpriseValueOverEBIT": 0,
					  "enterpriseValueOverEBITDA": 0,
					  "freeCashFlow": 0,
					  "freeCashFlowPerShare": 0,
					  "foreignCurrencyUSDExchangeRate": 0,
					  "grossProfit": 0,
					  "grossMargin": 0,
					  "goodwillAndIntangibleAssets": 0,
					  "interestExpense": 0,
					  "investedCapital": 0,
					  "investedCapitalAverage": 0,
					  "inventory": 0,
					  "investments": 0,
					  "investmentsCurrent": 0,
					  "investmentsNonCurrent": 0,
					  "totalLiabilities": 0,
					  "currentLiabilities": 0,
					  "liabilitiesNonCurrent": 0,
					  "marketCapitalization": 0,
					  "netCashFlow": 0,
					  "netCashFlowBusinessAcquisitionsDisposals": 0,
					  "issuanceEquityShares": 0,
					  "issuanceDebtSecurities": 0,
					  "paymentDividendsOtherCashDistributions": 0,
					  "netCashFlowFromFinancing": 0,
					  "netCashFlowFromInvesting": 0,
					  "netCashFlowInvestmentAcquisitionsDisposals": 0,
					  "netCashFlowFromOperations": 0,
					  "effectOfExchangeRateChangesOnCash": 0,
					  "netIncome": 0,
					  "netIncomeCommonStock": 0,
					  "netIncomeCommonStockUSD": 0,
					  "netLossIncomeFromDiscontinuedOperations": 0,
					  "netIncomeToNonControllingInterests": 0,
					  "profitMargin": 0,
					  "operatingExpenses": 0,
					  "operatingIncome": 0,
					  "tradeAndNonTradePayables": 0,
					  "payoutRatio": 0,
					  "priceToBookValue": 0,
					  "priceEarnings": 0,
					  "priceToEarningsRatio": 0,
					  "propertyPlantEquipmentNet": 0,
					  "preferredDividendsIncomeStatementImpact": 0,
					  "sharePriceAdjustedClose": 0,
					  "priceSales": 0,
					  "priceToSalesRatio": 0,
					  "tradeAndNonTradeReceivables": 0,
					  "accumulatedRetainedEarningsDeficit": 0,
					  "revenues": 0,
					  "revenuesUSD": 0,
					  "researchAndDevelopmentExpense": 0,
					  "returnOnAverageAssets": 0,
					  "returnOnAverageEquity": 0,
					  "returnOnInvestedCapital": 0,
					  "returnOnSales": 0,
					  "shareBasedCompensation": 0,
					  "sellingGeneralAndAdministrativeExpense": 0,
					  "shareFactor": 0,
					  "shares": 0,
					  "weightedAverageShares": 0,
					  "weightedAverageSharesDiluted": 0,
					  "salesPerShare": 0,
					  "tangibleAssetValue": 0,
					  "taxAssets": 0,
					  "incomeTaxExpense": 0,
					  "taxLiabilities": 0,
					  "tangibleAssetsBookValuePerShare": 0,
					  "workingCapital": 0
					}
				  ]
				}
			`))
	}
}

func GetMarketStatus(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(`
			{
			  "market": "open",
			  "serverTime": "2018-07-19T08:51:07-04:00",
			  "exchanges": {
				"nyse": "open",
				"nasdaq": "open",
				"otc": "extended-hours"
			  },
			  "currencies": {
				"fx": "open",
				"crypto": "open"
			  }
			}
			`))
	}
}

func GetMarketStatusUpcoming(writer http.ResponseWriter, request *http.Request) {
	returnError := request.URL.Query().Get("RETURN_ERROR")
	if returnError != "" {
		writer.WriteHeader(403)
		r := make(map[string]string)
		r["status"] = "ERROR"
		r["request_id"] = "some request id"
		r["error"] = "You are not entitled to this data"
		b, e := json.Marshal(r)
		if e != nil {
			return
		}
		_, _ = writer.Write(b)

	} else {
		_, _ = writer.Write([]byte(`
				[
				  {
					"exchange": "NYSE",
					"name": "Thanksgiving Day",
					"status": "early-close",
					"date": "2018-11-23T00:00:00.000Z",
					"open": "2018-11-23T09:30:00.000Z",
					"close": "2018-11-23T13:00:00.000Z"
				  }
				]
			`))
	}
}
