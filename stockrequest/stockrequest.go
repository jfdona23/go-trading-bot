package stockrequest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jfdona23/go-trading-bot/helpers"
)

var log = helpers.GetLogger()

const alphaVantageBaseUrl string = "https://www.alphavantage.co/query?function="

type Stockrequest struct {
	token string
}

func New() Stockrequest {
	alphaVantageToken := os.Getenv("STOCK_TOKEN")
	sa := Stockrequest{alphaVantageToken}
	return sa
}

func New2() *Stockrequest {
	alphaVantageToken := os.Getenv("STOCK_TOKEN")
	return &Stockrequest{alphaVantageToken}
}

func httpGet(url string, log helpers.LoggerType) (*http.Response, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	response, err := httpClient.Get(url)
	if err != nil {
		return response, err
	}
	log.Trace("Request URL: ", response.Request.URL)
	log.Trace(fmt.Sprintf("Full Response: %+v", response))
	return response, nil
}

func httpGetJson(url string, target interface{}, log helpers.LoggerType) error {
	response, err := httpGet(url, log)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(target)
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("Parsed response: %+v", target))
	return nil
}

func (sa Stockrequest) GetSymbolGlobalQuote(symbol string) *GlobalQuoteResponse {
	url := alphaVantageBaseUrl + "GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + sa.token
	target := new(GlobalQuoteResponse)
	httpGetJson(url, target, log)
	if target.GQ.Symbol == "" {
		target.Error = "There was an error parsing the API response"
		log.Error("There was an error parsing the API response. Enable debug or trace to see more details.")
	}
	return target
}

func (sa Stockrequest) SearchSymbol(symbol string) *SearchSymbolResponse {
	url := alphaVantageBaseUrl + "SYMBOL_SEARCH&keywords=" + symbol + "&apikey=" + sa.token
	target := new(SearchSymbolResponse)
	httpGetJson(url, target, log)
	return target
}

func (sa Stockrequest) GetForex(fromCurrency string, toCurrency string) *ForexResponse {
	url := alphaVantageBaseUrl + "CURRENCY_EXCHANGE_RATE&from_currency=" + fromCurrency + "&to_currency=" + toCurrency + "&apikey=" + sa.token
	target := new(ForexResponse)
	httpGetJson(url, target, log)
	if target.FX.FromCurrencyCode == "" {
		target.Error = "There was an error parsing the API response"
		log.Error("There was an error parsing the API response. Enable debug or trace to see more details.")
	}
	return target
}

func (sa Stockrequest) GetCryptoRating(symbol string) *CryptoRatingResponse {
	url := alphaVantageBaseUrl + "CRYPTO_RATING&symbol=" + symbol + "&apikey=" + sa.token
	target := new(CryptoRatingResponse)
	httpGetJson(url, target, log)
	if target.CR.Symbol == "" {
		target.Error = "There was an error parsing the API response"
		log.Error("There was an error parsing the API response. Enable debug or trace to see more details.")
	}
	return target
}
