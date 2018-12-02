package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/logging"
)

type showProfileCommandHandler struct {
}

var showProfileCommandHandlerInstance *showProfileCommandHandler

func init() {
	showProfileCommandHandlerInstance = &showProfileCommandHandler{}
}

func (h *showProfileCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	logging.Info().Msg("some message has been created")
}

func CreateShowProfileCommandHandler() MessageCreatedHandler {
	return showProfileCommandHandlerInstance
}
