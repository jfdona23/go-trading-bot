# Go Trading Bot

Simple **Discord** bot to retrieve data about different stock symbols.

1. [Launching](#launching)
1. [Usage](#usage)
## Launching
It can be configured using environment variables as follows:

|   Variable  |       Description       |   Possible Values  | Default |
|:-----------:|:-----------------------:|:------------------:|:-------:|
| LOG         | Set the log level       | debug, error, info | info    |
| BOT_TOKEN   | Discord bot API token   |                    |         |
| STOCK_TOKEN | Alpha Vantage API token |                    |         |

Compile...
```go
go build -o mybot
```
...and launch:
```shell
LOG=error SOTCK_TOKEN=GSDRFLH3456T4L5 BOT_TOKEN=7V67SDRFG7TYERG34.VE45235TT.G5Y3546Y5 ./mybot
```

## Usage
| Command | Parameter |                  Description                  |      Example      |
|:-------:|:---------:|:---------------------------------------------:|:-----------------:|
| search  | <keyword> | Search possible symbols for the given keyword | search salesforce |
| global  | <symbol>  | Return the current price for the given symbol | search crm        |