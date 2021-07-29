package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Command line flags
var (
	level             = flag.String("l", "info", "Log level to use")
	authToken         = flag.String("t", "", "Bot authentication token")
	alphaVantageToken = flag.String("alphavantage", "", "Alpha Vantage API token")
	guildID           = flag.String("g", "", "Test guild ID. If not passed, bot registers commands globally")
)

func init() {
	flag.Parse()

	l := new(Logger)
	l.SetLevel(*level)
	SetLogger(l)
	Log.Info("Log level set to " + Log.GetLevel())

	if *authToken == "" {
		panic("no authenticatin token was provided")
	}

	if *alphaVantageToken == "" {
		Log.Warn("Alpha Vantage token not provided.")
	}
}

func main() {
	// Recover from panic to keep the bot uptime as high as possible
	defer catchPanic()

	bot := NewClient(*authToken, *guildID)

	session, err := bot.Start()
	if err != nil {
		Log.Error("There was an error starting the bot: " + err.Error())
	}

	Log.Info("Press CTRL-C to exit")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	err = bot.Stop(session)
	if err != nil {
		Log.Error("There was an error stoping the bot: " + err.Error())
	}
}

// Recovers from panic
func catchPanic() {
	if r := recover(); r != nil {
		Log.Error("Recovered from panic " + fmt.Sprintf("%#v", r))
	}
}
