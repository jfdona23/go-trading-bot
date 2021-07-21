package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jfdona23/go-trading-bot/client"
	"github.com/jfdona23/go-trading-bot/engine"
	"github.com/jfdona23/go-trading-bot/helpers"
)

var (
	token    string = os.Getenv("BOT_TOKEN")
	logLevel string = helpers.Getenv("LOG", "info")
	log             = helpers.GetLogger(logLevel)
)

func main() {
	startupTimeBegin := time.Now()
	log.Info("Log level set to ", log.GetLevel())

	bot := client.New(token)
	session := bot.Start()
	defer bot.Stop(session)

	bot.Handler(session, engine.MessageCreate)
	startupTimeEnd := time.Now()
	startupTimeTotal := (startupTimeEnd.UnixNano() - startupTimeBegin.UnixNano()) / 1000000
	log.Debug("Bot started in ", startupTimeTotal, " miliseconds")

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
