package handler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

type YesNoHandler struct{}

func (h *YesNoHandler) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ask",
		Description: "Ask a question and the bot will answer yes or no.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "question",
				Description: "Your question",
				Required:    true,
			},
		},
	}
}

func (h *YesNoHandler) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Get the question from the interaction Data
	questionOption := i.ApplicationCommandData().Options[0]
	question := questionOption.StringValue()

	// Seed the random generator
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Generate yes or no answer
	response := "yes"
	if rng.Intn(2) == 1 {
		response = "no"
	}

	// Send the response to the user
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("For the question: %s\nThe great GoBot Answers: %s", question, response),
		},
	})
	if err != nil {
		log.Errorf("error responding to yesno: %v", err)
		return
	}
}
