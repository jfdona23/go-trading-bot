package client

import (
	"os"

	"github.com/jfdona23/go-trading-bot/helpers"

	"github.com/bwmarrin/discordgo"
)

var log = helpers.GetLogger()

// Type Alias
type Session = *discordgo.Session
type MsgEvent = *discordgo.MessageCreate

type client struct {
	token string
}

func New() client {
	tokenBot := os.Getenv("BOT_TOKEN")
	c := client{tokenBot}
	return c
}

func (c client) Start() (Session, error) {
	log.Info("Starting Bot...")

	log.Debug("Authenticating against Discord...")
	session, err := discordgo.New("Bot " + c.token)
	if err != nil {
		log.Error(err)
		return session, err
	}

	log.Debug("Opening WebSocket connection...")
	err = session.Open()
	if err != nil {
		log.Error(err)
		return session, err
	}

	// Set the Intents (Events) to listen to - Calculated with https://ziad87.net/intents/
	const myIntents discordgo.Intent = 4608 // GuildMessages, DirectMessages
	session.Identify.Intents = myIntents
	log.Debug("Intents set to ", myIntents)

	log.Info("Bot successfully started")
	return session, nil
}

func (c client) Stop(session Session) error {
	log.Info("Stopping Bot...")
	return session.Close()
}

func (c client) Handler(session Session, handler interface{}) {
	log.Debug("Adding new handler: ", handler)
	session.AddHandler(handler)
}
