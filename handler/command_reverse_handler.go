package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type ReverseHandler struct{}

func (h *ReverseHandler) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "reverse",
		Description: "Give me a sentence and I revert it for you.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "sentence",
				Description: "Your sentence",
				Required:    true,
			},
		},
	}
}

func (h *ReverseHandler) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Get the sentence from the interaction Data
	sentenceOption := i.ApplicationCommandData().Options[0]
	sentence := sentenceOption.StringValue()

	revertedSentence := reverseString(sentence)

	// Send the response to the user
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("The reverse of the sentence is: %s", revertedSentence),
		},
	})

	if err != nil {
		log.Errorf("error responding to reverse: %v", err)
		return
	}
}

func reverseString(str string) string {
	bytes := []byte(str)
	reversed := []byte{}

	for i := len(bytes) - 1; i >= 0; i-- {
			reversed = append(reversed, bytes[i])
	}
	return string(reversed)
}