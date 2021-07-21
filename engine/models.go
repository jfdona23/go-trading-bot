package engine

import (
	"strconv"
	"strings"

	"github.com/jfdona23/go-trading-bot/stockrequest"
)

func CommandNotFound(cmd string) string {
	return "Command unknown  :point_right:  \"**" + cmd + "**\""
}

func ShowHelp() string {
	return "```\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"| Command |          Parameter             |                         Description                                |      Example      |\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"| help    |                                | Show this help                                                     |                   |\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"| search  |          <keyword>             | Search possible symbols for the given keyword                      | search salesforce |\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"| global  |          <symbol>              | Return the current price for the given symbol                      | search crm        |\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"| forex   | <from_currency> <to_currency>  | Return the current exchange rates for the given currencies         | forex ars usd     |\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"| crypto  |       <crypto_symbol>          | Fundamental Crypto Asset Score (FCAS) for the given cryptocurrency | crypto btc        |\n" +
		"+---------+--------------------------------+--------------------------------------------------------------------+-------------------+\n" +
		"```"
}

func GlobalQuoteIM(r *stockrequest.GlobalQuoteResponse) string {
	var arrowEmoji string
	if r.Error != "" {
		return r.Error
	}
	if strings.HasPrefix(r.GQ.Change, "-") {
		arrowEmoji = ":chart_with_downwards_trend:"
	} else {
		arrowEmoji = ":chart_with_upwards_trend:"
	}
	return ":money_with_wings: **Global data for " + r.GQ.Symbol + "**\n" +
		"        • Current Price: " + r.GQ.Price + "\n" +
		"        • Tendency: " + r.GQ.Change + " " + arrowEmoji
}

func SearchSymbolIM(r *stockrequest.SearchSymbolResponse) string {
	if len(r.Results) == 0 {
		return "No results found :cry:"
	}
	head := ":mag_right: **" + strconv.Itoa(len(r.Results)) + " Results found**\n"
	body := ""
	for idx, result := range r.Results {
		body = body + strconv.Itoa(idx+1) + " Symbol: " + result.Symbol + "\n" +
			"    Name: " + result.Name + "\n" +
			"    Region: " + result.Region + "\n" +
			"    Currency: " + result.Currency + "\n\n"
	}
	return head + "```" + body + "```"
}

func ForexIM(r *stockrequest.ForexResponse) string {
	if r.Error != "" {
		return r.Error
	}
	return ":currency_exchange: **Current exchange values**\n" +
		"**" + r.FX.FromCurrencyCode + "** (" + r.FX.FromCurrencyName + ")" + "  :arrow_right:  " +
		"**" + r.FX.ToCurrencyCode + "** (" + r.FX.ToCurrencyName + ")\n" +
		"        • Sell Price: " + r.FX.SellPrice + "\n" +
		"        • Buy Price: " + r.FX.BuyPrice + "\n"
}

func CryptoRatingIM(r *stockrequest.CryptoRatingResponse) string {
	if r.Error != "" {
		return r.Error
	}
	return ":coin: **Cryptocurrency rating for: " + r.CR.Symbol + "** (" + r.CR.Name + ")\n" +
		"        • Score: " + r.CR.RatingScore + "/1000 (" + r.CR.Rating + ")\n" +
		"        • Maturity: " + r.CR.MaturityScore + "/1000"
}
