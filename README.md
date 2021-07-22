# Go Trading Bot

Simple **Discord** bot to retrieve data about different stock symbols.

1. [Launching](#launching)
1. [Usage](#usage)
## Launching
It can be configured using environment variables as follows:

|   Variable  |       Description              |   Possible Values         | Default |
|:-----------:|:------------------------------:|:-------------------------:|:-------:|
| LOG         | Set the log level              | debug, error, info, trace | info    |
| APPID       | Set the Discord Application ID |                           |         |
| BOT_TOKEN   | Discord bot API token          |                           |         |
| STOCK_TOKEN | Alpha Vantage API token        |                           |         |

Compile...
```shell
go build -o mybot
```
...and launch:
```shell
APPID=123451234512345 LOG=error SOTCK_TOKEN=GSDRFLH3456T4L5 BOT_TOKEN=7V67SDRFG7TYERG34.VE45235TT.G5Y3546Y5 ./mybot
```

## Usage
You can summon the bot by chatting directly with it (DM) or by mentioning it (@botname) and the passing one of the following commands:

| Command | Parameter                          | Description                                                        | Example           |
|:-------:|:----------------------------------:|:------------------------------------------------------------------:|:-----------------:|
| help    |                                    | Show this help                                                     |                   |
| search  | \<keyword\>                        | Search possible symbols for the given keyword                      | search salesforce |
| global  | \<symbol\>                         | Return the current price for the given symbol                      | search crm        |
| forex   | \<from_currency\> \<to_currency\>  | Return the current exchange rates for the given currencies         | forex ars usd     |
| crypto  | \<crypto_symbol\>                  | Fundamental Crypto Asset Score (FCAS) for the given cryptocurrency | crypto btc        |