package main

type SearchSymbolResponse struct {
	Results *[]searchSymbolResult `json:"bestMatches"`
}

type searchSymbolResult struct {
	Symbol   string `json:"1. symbol"`
	Name     string `json:"2. name"`
	Region   string `json:"4. region"`
	Currency string `json:"8. currency"`
}

type StockPriceResponse struct {
	StockPrice *stockPriceContent `json:"Global Quote"`
}

type stockPriceContent struct {
	Symbol        string `json:"01. symbol"`
	PriceOpen     string `json:"02. open"`
	PriceHigh     string `json:"03. high"`
	PriceLow      string `json:"04. low"`
	Price         string `json:"05. price"`
	LastClose     string `json:"07. latest trading day"`
	PricePrevious string `json:"08. previous close"`
	Change        string `json:"10. change percent"`
}

type ForexResponse struct {
	Forex *forexContent `json:"Realtime Currency Exchange Rate"`
}

type forexContent struct {
	FromCurrencyCode string `json:"1. From_Currency Code"`
	FromCurrencyName string `json:"2. From_Currency Name"`
	ToCurrencyCode   string `json:"3. To_Currency Code"`
	ToCurrencyName   string `json:"4. To_Currency Name"`
	SellPrice        string `json:"8. Bid Price"`
	BuyPrice         string `json:"9. Ask Price"`
}

type CryptoRatingResponse struct {
	CryptoRating *cryptoRatingContent `json:"Crypto Rating (FCAS)"`
}

type cryptoRatingContent struct {
	Symbol        string `json:"1. symbol"`
	Name          string `json:"2. name"`
	Rating        string `json:"3. fcas rating"`
	RatingScore   string `json:"4. fcas score"`
	MaturityScore string `json:"6. market maturity score"`
}
