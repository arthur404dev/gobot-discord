package main

import (
	"github.com/arthur404dev/gobot-discord/config"
	"github.com/arthur404dev/gobot-discord/discord"
)

func main() {
	// Initialize the service Config
	config.Init()

	log := config.NewLogger("main")

	// Initialize the discord Bot

	discord.Init()

	log.Fatalf("no service to run. Exiting...")
}
