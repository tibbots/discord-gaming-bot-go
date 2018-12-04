package handler

import (
	"github.com/TurtleGamingFTW/dblgo-archive"
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/repository"
	"strconv"
)

type versionCommandHandler struct {
	versionMessage   *discordgo.MessageEmbed
	serverRepository repository.ServerRepository
}

func (h *versionCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || (!event.isCommand("version")) {
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.versionMessage)

	if !event.isFromAdmin() {
		return
	}

	servers, err := h.serverRepository.GetAll()
	if err != nil {
		event.LogError(err, "version command failed displaying stats")
		return
	}

	channels := 0
	members := 0
	active := 0
	inactive := 0
	for _, server := range servers {
		if server.Deleted != 0 {
			inactive++
			continue
		}
		active++
		channels += server.Channels
		members += server.Members
	}

	botIdAsInt, err := strconv.Atoi(s.State.User.ID)
	if err != nil {
		event.LogError(err, "unable to convert bot id to integer "+s.State.User.ID)
		return
	}
	discordBotsClient := dblgo.NewDBL(environment.GetEnvironment().DiscordBotsToken, s.State.User.ID)
	dblServerStats, err := discordBotsClient.GetStats(botIdAsInt)
	if err != nil {
		event.LogError(err, "unable to convert bot id to integer "+s.State.User.ID)
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Bot Stats",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
		},

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Active Servers",
				Value:  strconv.Itoa(active),
				Inline: true,
			},
			{
				Name:   "Overall Installs",
				Value:  strconv.Itoa(active + inactive),
				Inline: true,
			},
			{
				Name:   "Active Members",
				Value:  strconv.Itoa(members),
				Inline: true,
			},
			{
				Name:   "Active Channels",
				Value:  strconv.Itoa(channels),
				Inline: true,
			},
			{
				Name:   "Server-Stats on discordbots.org",
				Value:  strconv.Itoa(dblServerStats.ServerCount),
				Inline: false,
			},
		},
	})

}

func CreateVersionCommandHandler(serverRepository repository.ServerRepository) MessageCreatedHandler {
	return &versionCommandHandler{
		serverRepository: serverRepository,
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
