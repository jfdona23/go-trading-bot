package engine

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jfdona23/go-trading-bot/client"
	"github.com/jfdona23/go-trading-bot/helpers"
	"github.com/jfdona23/go-trading-bot/stockrequest"
)

var (
	botUser          string = "<@866098491121598526>"
	botNickname      string = "<@!866098491121598526>"
	logLevel         string = helpers.Getenv("LOG", "info")
	log                     = helpers.GetLogger(logLevel)
	checkErrorAndLog        = helpers.CheckErrorAndLog
)

func MessageCreate(s client.Session, m client.MsgEvent) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// For mentions
	if strings.HasPrefix(m.Content, botUser) || strings.HasPrefix(m.Content, botNickname) {
		var msgTrimmed string
		msgTrimmed = strings.TrimPrefix(m.Content, botUser+" ")
		msgTrimmed = strings.TrimPrefix(msgTrimmed, botNickname+" ")

		message := parseMessage(msgTrimmed)
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			log.Error(err)
		}
		return
	}

	// For Direct Messages
	isDM, err := isDirectMessage(s, m)
	if err != nil {
		log.Error(err)
		return
	}

	if isDM {
		message := parseMessage(m.Content)
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		checkErrorAndLog(err)
		return
	}
}

func isDirectMessage(s client.Session, m client.MsgEvent) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}

func parseMessage(message string) string {
	msgArray := strings.Split(message, " ")
	cmd := strings.ToLower(msgArray[0])

	switch cmd {
	case "help":
		return ShowHelp()
	case "global":
		response := stockrequest.GetSymbolGlobalQuote(msgArray[1])
		return GlobalQuoteIM(response)
	case "search":
		response := stockrequest.SearchSymbol(msgArray[1])
		return SearchSymbolIM(response)
	case "forex":
		response := stockrequest.GetForex(msgArray[1], msgArray[2])
		return ForexIM(response)
	case "crypto":
		response := stockrequest.GetCryptoRating(msgArray[1])
		return CryptoRatingIM(response)
	default:
		return CommandNotFound(cmd)
	}
}
