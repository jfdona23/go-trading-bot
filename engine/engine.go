package engine

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jfdona23/go-trading-bot/client"
	"github.com/jfdona23/go-trading-bot/helpers"
	"github.com/jfdona23/go-trading-bot/stockrequest"
)

var log = helpers.GetLogger()

type Engine struct {
	botUser     string
	botNickname string
	stock       stockrequest.Stockrequest
}

func New(stock stockrequest.Stockrequest) Engine {
	var (
		appID       string = os.Getenv("APPID")
		botUser     string = "<@" + appID + ">"
		botNickname string = "<@!" + appID + ">"
	)
	e := Engine{botUser, botNickname, stock}
	return e
}

func (e Engine) MessageCreate(s client.Session, m client.MsgEvent) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// For mentions
	if strings.HasPrefix(m.Content, e.botUser) || strings.HasPrefix(m.Content, e.botNickname) {
		var msgTrimmed string
		msgTrimmed = strings.TrimPrefix(m.Content, e.botUser+" ")
		msgTrimmed = strings.TrimPrefix(msgTrimmed, e.botNickname+" ")

		message := parseMessage(e.stock, msgTrimmed)
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
		message := parseMessage(e.stock, m.Content)
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			log.Error(err)
		}
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

func parseMessage(stock stockrequest.Stockrequest, message string) string {
	msgArray := strings.Split(message, " ")
	cmd := strings.ToLower(msgArray[0])

	switch cmd {
	case "help":
		return ShowHelp()
	case "global":
		response := stock.GetSymbolGlobalQuote(msgArray[1])
		return GlobalQuoteIM(response)
	case "search":
		response := stock.SearchSymbol(msgArray[1])
		return SearchSymbolIM(response)
	case "forex":
		response := stock.GetForex(msgArray[1], msgArray[2])
		return ForexIM(response)
	case "crypto":
		response := stock.GetCryptoRating(msgArray[1])
		return CryptoRatingIM(response)
	default:
		return CommandNotFound(cmd)
	}
}
