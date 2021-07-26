package main

import "github.com/bwmarrin/discordgo"

// Shorthand for slash command names
var (
	searchSymbol string = "search-symbol"
	stockPrice   string = "stock-price"
	forex        string = "forex"
	cryptoRating string = "crypto-rating"
)

// Bot Slash commands definitions
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        searchSymbol,
		Description: "Searches for a sotck symbol",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "keyword",
				Description: "Keyword for searching",
				Required:    true,
			},
		},
	},
	{
		Name:        stockPrice,
		Description: "Shows current price and tendency for a given symbol",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "symbol",
				Description: "Symbol to show the price for",
				Required:    true,
			},
		},
	},
	{
		Name:        forex,
		Description: "Shows current exchange rates for the given currencies",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "from-currency",
				Description: "Source currency",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "to-currency",
				Description: "Destination currency",
				Required:    true,
			},
		},
	},
	{
		Name:        cryptoRating,
		Description: "Shows Fundamental Crypto Asset Score (FCAS) for a given crypto-currency",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "crypto-symbol",
				Description: "Symbol to show the FCAS for",
				Required:    true,
			},
		},
	},
}

// Bot Slash commands handlers
var commandsHandlers = map[string]func(s Session, i Interaction){
	searchSymbol: searchSymbolHandler,
	stockPrice:   stockPriceHandler,
	forex:        forexHandler,
	cryptoRating: cryptoRatingHandler,
}
