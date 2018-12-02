package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/logging"
)

type addAccountHandler struct {
}

func (h *addAccountHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	logging.Info().Msg("some message has been created")
}

func GetAddAccountHandler() MessageCreatedHandler {
	return &addAccountHandler{}
}
