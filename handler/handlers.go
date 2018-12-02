package handler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type MessageCreatedHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
}

type MessageCreatedEvent struct {
	state        *discordgo.Session
	message      *discordgo.MessageCreate
	plainContent string
}

func GetMessageCreatedEvent(s *discordgo.Session, m *discordgo.MessageCreate) *MessageCreatedEvent {
	return &MessageCreatedEvent{
		state:        s,
		message:      m,
		plainContent: "lazy",
	}
}

func (m *MessageCreatedEvent) GetPlainContent() string {
	if m.plainContent == "lazy" {
		m.plainContent = m.message.Content
		for _, user := range m.message.Mentions {
			m.plainContent = strings.TrimSpace(strings.NewReplacer(
				"<@"+user.ID+">", "",
				"<@!"+user.ID+">", "",
			).Replace(m.plainContent))
		}
	}

	return m.plainContent
}

func (m *MessageCreatedEvent) isSentByMe() bool {
	return m.message.ID == m.state.State.User.ID
}

func (m *MessageCreatedEvent) isSentByBot() bool {
	return m.message.Author.Bot
}

func (m *MessageCreatedEvent) isTalkingToMe() bool {
	return m.isDirectMessage() || m.isMentioningMe()
}

func (m *MessageCreatedEvent) isDirectMessage() bool {
	channel, err := m.state.Channel(m.message.ChannelID)
	if err != nil {
		return false
	}
	return channel.Type == discordgo.ChannelTypeDM
}

func (m *MessageCreatedEvent) isMentioningMe() bool {
	for _, mention := range m.message.Mentions {
		if mention.ID == m.state.State.User.ID {
			return true
		}
	}
	return false
}

func (m *MessageCreatedEvent) isCommand(command string) bool {
	return strings.HasPrefix(m.GetPlainContent(), command)
}

func (m *MessageCreatedEvent) shouldBeHandled() bool {
	return !m.isSentByMe() && !m.isSentByBot()
}
