package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"strings"
)

type MessageCreatedHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
}

type MessageCreatedEvent struct {
	state        *discordgo.Session
	message      *discordgo.MessageCreate
	plainContent string
	errorMessage *discordgo.MessageEmbed
}

func GetMessageCreatedEvent(s *discordgo.Session, m *discordgo.MessageCreate) *MessageCreatedEvent {
	return &MessageCreatedEvent{
		state:        s,
		message:      m,
		plainContent: "lazy",
		errorMessage: &discordgo.MessageEmbed{
			Title:       "Oops, something went horribly wrong on my side!",
			Description: "Stay tuned, our developers will have a look at it.",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},
		},
	}
}

func (m *MessageCreatedEvent) LogError(err error, message string) {
	logging.Error().
		Err(err).
		Msg(message)
}

func (m *MessageCreatedEvent) LogErrorAndRespond(err error, message string) {
	m.LogError(err, message)
	_, err = m.state.ChannelMessageSendEmbed(m.message.ChannelID, m.errorMessage)
	if err != nil {
		m.LogError(err, message)
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

func (m *MessageCreatedEvent) GetTargetedUser() *discordgo.User {
	for _, mention := range m.message.Mentions {
		if mention.ID == m.state.State.User.ID {
			continue
		}
		return mention
	}
	return m.message.Author
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

func (m *MessageCreatedEvent) isFromAdmin() bool {
	return m.message.Author.ID == "171919271272644609"
}

func (m *MessageCreatedEvent) shouldBeHandled() bool {
	return !m.isSentByMe() && !m.isSentByBot()
}
