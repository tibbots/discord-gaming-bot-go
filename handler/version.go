package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/environment"
)

type versionCommandHandler struct {
	versionMessage *discordgo.MessageEmbed
}

func (h *versionCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || (!event.isCommand("version")) {
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.versionMessage)

}

func CreateVersionCommandHandler() MessageCreatedHandler {
	return &versionCommandHandler{
		versionMessage: &discordgo.MessageEmbed{
			Title: "Bot Details",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Version",
					Value: environment.GetEnvironment().ProjectVersion,
				},
				{
					Name:  "Github Releases",
					Value: environment.GetEnvironment().ProjectUrl + "/releases",
				},
			},
		},
	}
}
