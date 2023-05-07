package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arthur404dev/gobot-discord/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Initialize the service Config
	config.Init()

	// Create a new Discord Session
	token := config.GetEnv().BOT_TOKEN
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord Session: ", err)
		return
	}

	// dg.AddHandler(messageCreate)
	// dg.Identify.Intents = discordgo.IntentGuildMessages

	dg.AddHandler(interactionCreate)

	// Abrir Conexao Websocket
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection with discord: ", err)
		return
	}

	err = createSlashCommand(dg, dg.State.User.ID)
	if err != nil {
		fmt.Println("Error Creating Slash Commands: ", err)
		return
	}

	fmt.Println("GoBot is running. Press Ctrl + C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "oibot" {
		s.ChannelMessageSend(m.ChannelID, "Oi Meu lindo maravilhoso!")
	}
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "oibot":
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Olá Seu Lindo, Bom te ver por aqui!",
				},
			})
			if err != nil {
				fmt.Println("Error Greeting User: ", err)
			}
		}
	}
}

func createSlashCommand(s *discordgo.Session, applicationID string) error {
	oiCommand := &discordgo.ApplicationCommand{
		Name:        "oibot",
		Description: "Cumprimenta o Bot e ele te Cumprimenta de volta!",
	}

	_, err := s.ApplicationCommandCreate(applicationID, "", oiCommand)
	if err != nil {
		return fmt.Errorf("error creating oiCommand")
	}
	return nil
}
