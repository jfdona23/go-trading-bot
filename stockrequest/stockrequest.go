package stockrequest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jfdona23/go-trading-bot/helpers"
)

var (
	token            string = os.Getenv("STOCK_TOKEN")
	logLevel         string = helpers.Getenv("LOG", "info")
	log                     = helpers.GetLogger(logLevel)
	checkErrorAndLog        = helpers.CheckErrorAndLog
)

func httpGet(url string) []byte {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	response, err := httpClient.Get(url)
	checkErrorAndLog(err)
	extractedBody := extractBody(response, 4096)
	log.Trace("Request URL: ", response.Request.URL)
	log.Trace(fmt.Sprintf("Full Response: %+v", response))
	log.Trace("Plain Response body: ", string(extractedBody))
	return extractedBody
}

func httpGetJson(url string, target interface{}) error {
	response := httpGet(url)
	// defer response.Body.Close()
	// err := json.NewDecoder(response.Body).Decode(target)
	err := json.Unmarshal(response, target)
	checkErrorAndLog(err)
	log.Debug(fmt.Sprintf("Parsed response: %+v", target))
	return err
}

func extractBody(response *http.Response, maxBytes int64) []byte {
	body, err := io.ReadAll(io.LimitReader(response.Body, maxBytes))
	checkErrorAndLog(err)
	// This is for response re-utilization
	// response.Body = io.NopCloser(bytes.NewBuffer(body))
	return body
}

func GetSymbolGlobalQuote(symbol string) *GlobalQuoteResponse {
	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + token
	target := new(GlobalQuoteResponse)
	httpGetJson(url, target)
	if target.GQ.Symbol == "" {
		target.Error = "There was an error parsing the API response"
		log.Error("There was an error parsing the API response. Enable debug or trace to see more details.")
	}
	return target
}

func SearchSymbol(symbol string) *SearchSymbolResponse {
	url := "https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=" + symbol + "&apikey=" + token
	target := new(SearchSymbolResponse)
	httpGetJson(url, target)
	return target
}

func GetForex(fromCurrency string, toCurrency string) *ForexResponse {
	url := "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=" + fromCurrency + "&to_currency=" + toCurrency + "&apikey=" + token
	target := new(ForexResponse)
	httpGetJson(url, target)
	if target.FX.FromCurrencyCode == "" {
		target.Error = "There was an error parsing the API response"
		log.Error("There was an error parsing the API response. Enable debug or trace to see more details.")
	}
	return target
}

func GetCryptoRating(symbol string) *CryptoRatingResponse {
	url := "https://www.alphavantage.co/query?function=CRYPTO_RATING&symbol=" + symbol + "&apikey=" + token
	target := new(CryptoRatingResponse)
	httpGetJson(url, target)
	if target.CR.Symbol == "" {
		target.Error = "There was an error parsing the API response"
		log.Error("There was an error parsing the API response. Enable debug or trace to see more details.")
	}
	return target
}
