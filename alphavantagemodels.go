package main

// Struct to unmarshal a JSON response comming from the SYMBOL_SEARCH endpoint (Alpha Vantage)
type SearchSymbolResponse struct {
	Results *[]SearchSymbolResult `json:"bestMatches"`
}

type SearchSymbolResult struct {
	Symbol   string `json:"1. symbol"`
	Name     string `json:"2. name"`
	Region   string `json:"4. region"`
	Currency string `json:"8. currency"`
}

// Struct to unmarshal a JSON response comming from the GLOBAL_QUOTE endpoint (Alpha Vantage)
type StockPriceResponse struct {
	StockPrice *StockPriceContent `json:"Global Quote"`
}

type StockPriceContent struct {
	Symbol        string `json:"01. symbol"`
	PriceOpen     string `json:"02. open"`
	PriceHigh     string `json:"03. high"`
	PriceLow      string `json:"04. low"`
	Price         string `json:"05. price"`
	LastClose     string `json:"07. latest trading day"`
	PricePrevious string `json:"08. previous close"`
	Change        string `json:"10. change percent"`
}

// Struct to unmarshal a JSON response comming from the CURRENCY_EXCHANGE_RATE endpoint (Alpha Vantage)
type ForexResponse struct {
	Forex *ForexContent `json:"Realtime Currency Exchange Rate"`
}

type ForexContent struct {
	FromCurrencyCode string `json:"1. From_Currency Code"`
	FromCurrencyName string `json:"2. From_Currency Name"`
	ToCurrencyCode   string `json:"3. To_Currency Code"`
	ToCurrencyName   string `json:"4. To_Currency Name"`
	SellPrice        string `json:"8. Bid Price"`
	BuyPrice         string `json:"9. Ask Price"`
}

// Struct to unmarshal a JSON response comming from the CRYPTO_RATING endpoint (Alpha Vantage)
type CryptoRatingResponse struct {
	CryptoRating *CryptoRatingContent `json:"Crypto Rating (FCAS)"`
}

type CryptoRatingContent struct {
	Symbol        string `json:"1. symbol"`
	Name          string `json:"2. name"`
	Rating        string `json:"3. fcas rating"`
	RatingScore   string `json:"4. fcas score"`
	MaturityScore string `json:"6. market maturity score"`
}
