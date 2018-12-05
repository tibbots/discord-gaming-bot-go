package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/help"
)

type helpCommandHandler struct {
}

func (h *helpCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || (!event.isCommand("help") && !event.isCommand("rtfm")) {
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, help.GetMessages().Help)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, help.GetMessages().SupportedPlatforms)
	_, _ = s.ChannelMessageSend(m.ChannelID, help.GetMessages().LinkPreviewHint)

}

func CreateHelpCommandHandler() MessageCreatedHandler {
	return &helpCommandHandler{}
}
