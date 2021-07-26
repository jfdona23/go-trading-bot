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
	switch typ := r.(type) {
	case *SearchSymbolResponse:
		return !reflect.DeepEqual(typ, new(SearchSymbolResponse))
	case *StockPriceResponse:
		return !reflect.DeepEqual(typ, new(StockPriceResponse))
	case *ForexResponse:
		return !reflect.DeepEqual(typ, new(ForexResponse))
	case *CryptoRatingResponse:
		return !reflect.DeepEqual(typ, new(CryptoRatingResponse))
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
