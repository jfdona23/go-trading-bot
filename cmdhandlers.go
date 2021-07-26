package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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

func searchSymbolHandler(s Session, i Interaction) {
	keyword := i.ApplicationCommandData().Options[0].StringValue()
	response := searchSymbolRequest(keyword)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: searchSymbolBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

func stockPriceHandler(s Session, i Interaction) {
	symbol := i.ApplicationCommandData().Options[0].StringValue()
	response := stockPriceRequest(symbol)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: stockPriceBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

func forexHandler(s Session, i Interaction) {
	fromCurrency := i.ApplicationCommandData().Options[0].StringValue()
	toCurrency := i.ApplicationCommandData().Options[1].StringValue()
	response := forexRequest(fromCurrency, toCurrency)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: forexBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}

func cryptoRatingHandler(s Session, i Interaction) {
	symbol := i.ApplicationCommandData().Options[0].StringValue()
	response := cryptoRatingRequest(symbol)
	interactionResponseData := &discordgo.InteractionResponseData{
		Content: cryptoRatingBuildIM(response),
	}
	interactionResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: interactionResponseData,
	}
	s.InteractionRespond(i.Interaction, interactionResponse)
}
