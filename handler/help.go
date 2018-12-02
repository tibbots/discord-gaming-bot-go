package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/logging"
)

type helpCommandHandler struct {
}

var helpCommandHandlerInstance *helpCommandHandler

func init() {
	helpCommandHandlerInstance = &helpCommandHandler{}
}

func (h *helpCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	logging.Info().Msg("some message has been created")
}

func GetHelpCommandHandler() MessageCreatedHandler {
	return helpCommandHandlerInstance
}
