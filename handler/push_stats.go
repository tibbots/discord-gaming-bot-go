package handler

import (
	"fmt"
	"github.com/TurtleGamingFTW/dblgo-archive"
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"github.com/tibbots/discord-gaming-bot-go/repository"
)

type pushDiscordBotsCommandHandler struct {
	serverRepository repository.ServerRepository
}

func (h *pushDiscordBotsCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.isFromAdmin() || !event.shouldBeHandled() || !event.isTalkingToMe() || !event.isCommand("push stats") {
		return
	}

	servers, err := h.serverRepository.GetAll()
	if err != nil {
		event.LogErrorAndRespond(err, "push command failed")
		return
	}

	discordBotsClient := dblgo.NewDBL(environment.GetEnvironment().DiscordBotsToken, s.State.User.ID)
	err = discordBotsClient.PostStats(len(servers))
	if err != nil {
		event.LogErrorAndRespond(err, "pushing discord-bots stats failed")
		return
	}
	logging.Info().
		Int("servers", len(servers)).
		Msg("pushed stats successfully to discordbots.org")

	err = discordBotsClient.PostStatsSharded(1, 0, 0)
	if err != nil {
		event.LogErrorAndRespond(err, "pushing discord-bots sharding stats failed")
		return
	}

	logging.Info().
		Int("shards", 1).
		Msg("pushed sharding stats successfully to discordbots.org")

	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("pushed %d servers and %d shards to discordbots.org", len(servers), 1))
}

func CreatePushStatsCommandHandler(serverRepository repository.ServerRepository) MessageCreatedHandler {
	return &pushDiscordBotsCommandHandler{
		serverRepository: serverRepository,
	}
}
