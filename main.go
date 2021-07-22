package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jfdona23/go-trading-bot/client"
	"github.com/jfdona23/go-trading-bot/engine"
	"github.com/jfdona23/go-trading-bot/helpers"
	"github.com/jfdona23/go-trading-bot/stockrequest"
)

var log = helpers.GetLogger()

func main() {
	startupTimeBegin := time.Now()
	log.Info("Log level set to ", log.GetLevel())

	// Start of DI
	stockrequest := stockrequest.New()
	engine := engine.New(stockrequest)
	bot := client.New()
	// End of DI

	session, err := bot.Start()
	if err != nil {
		log.Panic(err)
	}

	bot.Handler(session, engine.MessageCreate)

	startupTimeEnd := time.Now()
	startupTimeTotal := (startupTimeEnd.UnixNano() - startupTimeBegin.UnixNano()) / 1000000
	log.Debug("Bot started in ", startupTimeTotal, " miliseconds")

	log.Info("Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = bot.Stop(session)
	if err != nil {
		log.Panic(err)
	}
}
