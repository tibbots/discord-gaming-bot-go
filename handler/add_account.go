package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"strings"
)

type addAccountCommandHandler struct {
}

var addAccountCommandHandlerInstance *addAccountCommandHandler

func init() {
	addAccountCommandHandlerInstance = &addAccountCommandHandler{}
}

func (h *addAccountCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || !event.isCommand("add account") {
		return
	}

	fields := strings.Fields(event.GetPlainContent())
	if len(fields) != 2 {

		return
	}

	logging.Info().Msg("some message has been created")
}

func GetAddAccountCommandHandler() MessageCreatedHandler {
	return addAccountCommandHandlerInstance
}
