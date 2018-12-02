package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type helpCommandHandler struct {
	helpMessages *discordgo.MessageEmbed
}

func (h *helpCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || (!event.isCommand("help") && !event.isCommand("rtfm")) {
		return
	}

	if h.helpMessages == nil {
		h.helpMessages = &discordgo.MessageEmbed{
			Title: "Discord Gaming Bot Manual",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "made by tibbot.org",
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Creating a Profile",
					Value: fmt.Sprintf("@%s create profile (or via direct message)", s.State.User.Username),
				},
				{
					Name:  "Deleting a Profile",
					Value: fmt.Sprintf("@%s delete profile (or via direct message)", s.State.User.Username),
				},
				{
					Name:  "Inspecting your Profile",
					Value: fmt.Sprintf("@%s show profile (or via direct message)", s.State.User.Username),
				},
				{
					Name:  "Inspecting another persons Profile",
					Value: fmt.Sprintf("@%s show profile @YourFriend (or via direct message)", s.State.User.Username),
				}, {
					Name:  "Adding an platform account (like Steam, Origin or whatsoever)",
					Value: fmt.Sprintf("@%s add account [platform] [account-id] (or via direct message)", s.State.User.Username),
				},
			},
		}
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.helpMessages)

}

func GetHelpCommandHandler() MessageCreatedHandler {
	return &helpCommandHandler{}
}
