# Go Trading Bot

Simple **Discord** bot to retrieve data about different stock symbols.

1. [Launching](#launching)
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
