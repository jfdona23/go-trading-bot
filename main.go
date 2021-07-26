package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	level             = flag.String("l", "info", "Log level to use")
	authToken         = flag.String("t", "", "Bot authentication token")
	alphaVantageToken = flag.String("alphavantage", "", "Alpha Vantage API token")
	guildID           = flag.String("g", "", "Test guild ID. If not passed, bot registers commands globally")
)

func init() {
	flag.Parse()

	l := new(logger)
	l.setLevel(*level)
	setLogger(l)
	Log.Info("Log level set to " + Log.getLevel())

	if *authToken == "" {
		panic("no authenticatin token was provided")
	}

	if *alphaVantageToken == "" {
		Log.Warn("Alpha Vantage token not provided. Things may go wrong")
	}
}

func main() {
	bot := NewClient(*authToken, *guildID)

	session, err := bot.Start()
	if err != nil {
		Log.Error("There was an error starting the bot: " + err.Error())
	}

	Log.Info("Press CTRL-C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	err = bot.Stop(session)
	if err != nil {
		Log.Error("There was an error stoping the bot: " + err.Error())
	}
}
