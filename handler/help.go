package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/environment"
)

type helpCommandHandler struct {
	helpMessage *discordgo.MessageEmbed
}

func (h *helpCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || (!event.isCommand("help") && !event.isCommand("rtfm")) {
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.helpMessage)

}

func CreateHelpCommandHandler() MessageCreatedHandler {
	return &helpCommandHandler{
		helpMessage: &discordgo.MessageEmbed{
			Title: "Discord Gaming Bot Manual",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Creating a Profile",
					Value: "@MyNameOnYourServer create profile (or via direct message)",
				},
				{
					Name:  "Deleting a Profile",
					Value: "@MyNameOnYourServer delete profile (or via direct message)",
				},
				{
					Name:  "Inspecting your Profile",
					Value: "@MyNameOnYourServer show profile (or via direct message)",
				},
				{
					Name:  "Inspecting another persons Profile",
					Value: "@MyNameOnYourServer show profile @YourFriend (or via direct message)",
				}, {
					Name:  "Adding an platform account (like Steam, Origin or whatsoever)",
					Value: "@MyNameOnYourServer add account [platform] [account-id] (or via direct message)",
				},
			},
		},
	}
}
