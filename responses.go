package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	cmdNotFound string = "Command `%s` not found :disappointed:"
	errorMsg    string = ":cry: There was an error.\n"
)

// Validate if the response structure is properly fullfilled by comparing it with an empty one
func isValidResponseStruct(r interface{}) bool {
	switch r := r.(type) {
	// There is no invalid input for SearchSymbol, it just returns an empty slice of results.
	case *SearchSymbolResponse:
		return true
	// The Stock Price error is an empty struct, including nested ones.
	case *StockPriceResponse:
		empty := &StockPriceResponse{&stockPriceContent{}}
		return !reflect.DeepEqual(r, empty)
	// The Forex error is a totally different struct rather than an empty one.
	case *ForexResponse:
		empty := &ForexResponse{}
		return !reflect.DeepEqual(r, empty)
	// The Crypto Rating error is a totally different struct rather than an empty one.
	case *CryptoRatingResponse:
		empty := &CryptoRatingResponse{}
		return !reflect.DeepEqual(r, empty)
	}
	return false
}

// Build an Instant Message from a searchSymbolResponse object
func searchSymbolBuildIM(r *SearchSymbolResponse) string {
	if !isValidResponseStruct(r) {
		return errorMsg
	}
	if len(*r.Results) == 0 {
		return ":mag_right: **No results found**"
	}
	head := ":mag_right: **%d Results found**\n```%s```"
	body := ""
	for idx, result := range *r.Results {
		body = body + strconv.Itoa(idx+1) + " Symbol: " + result.Symbol + "\n" +
			"    Name: " + result.Name + "\n" +
			"    Region: " + result.Region + "\n" +
			"    Currency: " + result.Currency + "\n\n"
	}
	return fmt.Sprintf(head, len(*r.Results), body)
}

func stockPriceBuildIM(r *StockPriceResponse) string {
	if !isValidResponseStruct(r) {
		return errorMsg
	}
	var arrowEmoji string
	if strings.HasPrefix(r.StockPrice.Change, "-") {
		arrowEmoji = ":chart_with_downwards_trend:"
	} else {
		arrowEmoji = ":chart_with_upwards_trend:"
	}
	return ":money_with_wings: **Global data for " + r.StockPrice.Symbol + "**\n" +
		"        • Current Price: " + r.StockPrice.Price + "\n" +
		"        • Tendency: " + r.StockPrice.Change + " " + arrowEmoji
}

func forexBuildIM(r *ForexResponse) string {
	if !isValidResponseStruct(r) {
		return errorMsg
	}
	return ":currency_exchange: **Current exchange values**\n" +
		"**" + r.Forex.FromCurrencyCode + "** (" + r.Forex.FromCurrencyName + ")" + "  :arrow_right:  " +
		"**" + r.Forex.ToCurrencyCode + "** (" + r.Forex.ToCurrencyName + ")\n" +
		"        • Sell Price: " + r.Forex.SellPrice + "\n" +
		"        • Buy Price: " + r.Forex.BuyPrice + "\n"
}

func cryptoRatingBuildIM(r *CryptoRatingResponse) string {
	if !isValidResponseStruct(r) {
		return errorMsg
	}
	return ":coin: **Cryptocurrency rating for: " + r.CryptoRating.Symbol + "** (" + r.CryptoRating.Name + ")\n" +
		"        • Score: " + r.CryptoRating.RatingScore + "/1000 (" + r.CryptoRating.Rating + ")\n" +
		"        • Maturity: " + r.CryptoRating.MaturityScore + "/1000"
}
