package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/logging"
)

type deleteProfileCommandHandler struct {
}

var deleteProfileCommandHandlerInstance *deleteProfileCommandHandler

func init() {
	deleteProfileCommandHandlerInstance = &deleteProfileCommandHandler{}
}

func (h *deleteProfileCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	logging.Info().Msg("some message has been created")
}

func GetDeleteProfileCommandHandler() MessageCreatedHandler {
	return deleteProfileCommandHandlerInstance
}
