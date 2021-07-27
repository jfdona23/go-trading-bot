/*
Simple Discord bot to retrieve data about different stock symbols.

It can be configured using a simple set of flags as follows:

  -h, --help     Show this help.
  -l             Set the log level (default=info).
  -b             Discord bot API token.
  -g             GuildID for testing purposes. If not passed the bot works globally (default="").
  -alphavantage  Alpha Vantage API token.


Compile:
  go build -o mybot

Launch:
  ./mybot -l error -b 7V67SDRFG7TYERG34.VE45235TT.G5Y3546Y5 -alphavantage GSDRFLH3456T4L5


You can summon the bot using slash commands (`/cmd`) as follows:

  /search-symbol "keyword"                       Search possible symbols for the given keyword.
  /stock-price   "symbol"                        Return the current price for the given symbol.
  /forex         "from_currency" "to_currency"   Return the current exchange rates for the given currencies.
  /crypto-rating "crypto_symbol"                 Fundamental Crypto Asset Score (FCAS) for the given cryptocurrency.
*/
package main
