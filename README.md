# Go Trading Bot

Simple **Discord** bot to retrieve data about different stock symbols.

1. [Launching](#launching)
1. [Usage](#usage)
## Launching
It can be configured using a simple set of flags as follows:

|    Flag       |       Description                                                     |   Possible Values         | Default   |
|:-------------:|:---------------------------------------------------------------------:|:-------------------------:|:---------:|
| -h --help     | Show this help                                                        |                           |           |
| -l            | Set the log level                                                     | debug, error, info        | info      |
| -b            | Discord bot API token                                                 |                           |           |
| -g            | GuildID for testing purposes. If not passed the bot works globally.   |                           | \<blank\> |
| -alphavantage | Alpha Vantage API token                                               |                           |           |

Compile...
```shell
go build -o mybot
```
...and launch:
```shell
./mybot -l error -b 7V67SDRFG7TYERG34.VE45235TT.G5Y3546Y5 -alphavantage GSDRFLH3456T4L5 
```

## Usage
You can summon the bot using slash commands (`/cmd`) as follows:

| Command        | Parameter                          | Description                                                        |
|:--------------:|:----------------------------------:|:------------------------------------------------------------------:|
| search-symbol  | \<keyword\>                        | Search possible symbols for the given keyword                      |
| stock-price    | \<symbol\>                         | Return the current price for the given symbol                      |
| forex          | \<from_currency\> \<to_currency\>  | Return the current exchange rates for the given currencies         |
| crypto-rating  | \<crypto_symbol\>                  | Fundamental Crypto Asset Score (FCAS) for the given cryptocurrency |