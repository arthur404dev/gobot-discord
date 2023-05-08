package handler

import (
	"github.com/arthur404dev/gobot-discord/config"
)

var (
	env = config.GetEnv()
	log = config.NewLogger("handler")
)
