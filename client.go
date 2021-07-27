package main

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Type Aliases
type Session = *discordgo.Session
type Interaction = *discordgo.InteractionCreate

// The client struct
type Client struct {
	token string // Bot authentication token
	guild string // GuildID for testing purposes. If not passed, bot registers commands globally
}

// Constructor for a new Client struct
func NewClient(token string, guild string) Client {
	c := Client{token, guild}
	return c
}

// Start a new client session
func (c Client) Start() (Session, error) {
	Log.Info("Starting Bot")
	startupTimeBegin := time.Now()

	Log.Debug("Authenticating against Discord")
	session, err := discordgo.New("Bot " + c.token)
	if err != nil {
		return session, err
	}

	// Set the Intents (Events) to listen to - Calculated with https://ziad87.net/intents/
	const myIntents discordgo.Intent = 0 // Only listen to the always sent events
	session.Identify.Intents = myIntents
	Log.Debug("Intents set to " + strconv.Itoa(int(myIntents)))

	Log.Debug("Opening WebSocket connection")
	err = session.Open()
	if err != nil {
		return session, err
	}

	Log.Debug("Adding commands handlers")
	session.AddHandler(executeHandlerIfExists)

	Log.Debug("Publishing commands handlers")
	for _, appCmd := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, c.guild, appCmd)
		if err != nil {
			Log.Error("Can't create '" + appCmd.Name + "' command: " + err.Error())
		}
	}

	startupTimeEnd := time.Now()
	startupTimeTotal := (startupTimeEnd.UnixNano() - startupTimeBegin.UnixNano()) / 1000000
	Log.Info("Bot successfully started in " + strconv.FormatInt(startupTimeTotal, 10) + " miliseconds")

	return session, nil
}

// Close client's session
func (c Client) Stop(session Session) error {
	println("") // Print a blank line for aesthetical reasons.
	Log.Info("Stopping Bot")
	Log.Debug("Deleting slash commands")
	commands, err := session.ApplicationCommands(session.State.User.ID, c.guild)
	if err != nil {
		Log.Error("Error obtaining the list of commands to delete: " + err.Error())
	}
	for _, appCmd := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, c.guild, appCmd.ID)
		if err != nil {
			Log.Error("Can't delete '" + appCmd.Name + "' command: " + err.Error())
		}
	}
	return session.Close()
}

// If the command requested exists, execute its handler. Otherwise execute a command-not-found handler.
func executeHandlerIfExists(s Session, i Interaction) {
	command := i.ApplicationCommandData().Name
	if handler, isOK := commandsHandlers[command]; isOK {
		Log.Debug("Running handler for command: " + command)
		handler(s, i)
	} else {
		Log.Error("Handler not found for command: " + command)
		cmdNotFoundHandler(s, i)
	}
}
