package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/repository"
)

type trackBotJoinCommandHandler struct {
	serverRepository repository.ServerRepository
}

func (h *trackBotJoinCommandHandler) Handle(s *discordgo.Session, g *discordgo.GuildCreate) {
	_ = h.serverRepository.Persist(entity.CreateServer(g.Guild))
}

func CreateTrackBotJoinsCommand(serverRepository repository.ServerRepository) GuildCreateHandler {
	return &trackBotJoinCommandHandler{
		serverRepository: serverRepository,
	}
}
