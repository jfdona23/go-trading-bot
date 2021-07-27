package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Command not found, interaction handler
func cmdNotFoundHandler(s Session, i Interaction) {
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: fmt.Sprintf(cmdNotFound, i.ApplicationCommandData().Name),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

// SearchSymbolHandler, interaction handler
func SearchSymbolHandler(s Session, i Interaction) {
	keyword := i.ApplicationCommandData().Options[0].StringValue()
	response := SearchSymbolRequest(keyword)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: SearchSymbolBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

// StockPriceHandler, interaction handler
func StockPriceHandler(s Session, i Interaction) {
	symbol := i.ApplicationCommandData().Options[0].StringValue()
	response := StockPriceRequest(symbol)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: StockPriceBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

// ForexHandler, interaction handler
func ForexHandler(s Session, i Interaction) {
	fromCurrency := i.ApplicationCommandData().Options[0].StringValue()
	toCurrency := i.ApplicationCommandData().Options[1].StringValue()
	response := ForexRequest(fromCurrency, toCurrency)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: ForexBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

// CryptoRatingHandler, interaction handler
func CryptoRatingHandler(s Session, i Interaction) {
	symbol := i.ApplicationCommandData().Options[0].StringValue()
	response := CryptoRatingRequest(symbol)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: CryptoRatingBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}
