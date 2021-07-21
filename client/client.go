package client

import (
	"github.com/jfdona23/go-trading-bot/helpers"

	"github.com/bwmarrin/discordgo"
)

// Type Alias
type Session = *discordgo.Session
type MsgEvent = *discordgo.MessageCreate

type client struct {
	token string
}

var (
	logLevel           string = helpers.Getenv("LOG", "info")
	log                       = helpers.GetLogger(logLevel)
	checkErrorAndPanic        = helpers.CheckErrorAndPanic
)

func New(token string) client {
	c := client{token}
	return c
}

func (c client) Start() Session {
	log.Info("Starting Bot...")

	log.Debug("Authenticating against Discord...")
	session, err := discordgo.New("Bot " + c.token)
	checkErrorAndPanic(err)

	log.Debug("Opening WebSocket connection...")
	err = session.Open()
	checkErrorAndPanic(err)

	// Set the Intents where listen to
	// GuildMessages + DirectMessages - Calculated with https://ziad87.net/intents/
	const myIntents discordgo.Intent = 4608
	log.Debug("Set Intents to ", myIntents)
	session.Identify.Intents = myIntents

	log.Info("Bot successfully started")
	return session
}

func (c client) Stop(session Session) {
	log.Info("Stopping Bot...")
	session.Close()
}

func (c client) Handler(session Session, handler interface{}) {
	session.AddHandler(handler)
}
