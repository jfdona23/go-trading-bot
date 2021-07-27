package main

import "fmt"

// Alpha Vantage URLs and Endpoints
const (
	alphaVantageBaseUrl string = "https://www.alphavantage.co/query?function="
	searchSymbolUrl     string = alphaVantageBaseUrl + "SYMBOL_SEARCH&apikey=%s&keywords=%s"
	stockPriceUrl       string = alphaVantageBaseUrl + "GLOBAL_QUOTE&apikey=%s&symbol=%s"
	forexUrl            string = alphaVantageBaseUrl + "CURRENCY_EXCHANGE_RATE&apikey=%s&from_currency=%s&to_currency=%s"
	cryptoRatingUrl     string = alphaVantageBaseUrl + "CRYPTO_RATING&apikey=%s&symbol=%s"
)

// Function to look for stock symbols from a keyword
func SearchSymbolRequest(keyword string) *SearchSymbolResponse {
	url := fmt.Sprintf(searchSymbolUrl, *alphaVantageToken, keyword)
	response, err := HttpGet(url)
	if err != nil {
		Log.Error(err)
	}
	jsonResponse := new(SearchSymbolResponse)
	err = ResponseToJsonStruct(response, jsonResponse)
	if err != nil {
		Log.Error(err)
	}
	return jsonResponse
}

// Function to look for a symbol current trading data
func StockPriceRequest(symbol string) *StockPriceResponse {
	url := fmt.Sprintf(stockPriceUrl, *alphaVantageToken, symbol)
	response, err := HttpGet(url)
	if err != nil {
		Log.Error(err)
	}
	jsonResponse := new(StockPriceResponse)
	err = ResponseToJsonStruct(response, jsonResponse)
	if err != nil {
		Log.Error(err)
	}
	return jsonResponse
}

// Function to retrieve currency exchange rates
func ForexRequest(fromCurrency string, toCurrency string) *ForexResponse {
	url := fmt.Sprintf(forexUrl, *alphaVantageToken, fromCurrency, toCurrency)
	response, err := HttpGet(url)
	if err != nil {
		Log.Error(err)
	}
	jsonResponse := new(ForexResponse)
	err = ResponseToJsonStruct(response, jsonResponse)
	if err != nil {
		Log.Error(err)
	}
	return jsonResponse
}

// Function to retrieve crypto currency ratings (FCAS)
func CryptoRatingRequest(symbol string) *CryptoRatingResponse {
	url := fmt.Sprintf(cryptoRatingUrl, *alphaVantageToken, symbol)
	response, err := HttpGet(url)
	if err != nil {
		Log.Error(err)
	}
	jsonResponse := new(CryptoRatingResponse)
	err = ResponseToJsonStruct(response, jsonResponse)
	if err != nil {
		Log.Error(err)
	}
	return jsonResponse
}
