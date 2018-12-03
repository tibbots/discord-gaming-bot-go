package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/repository"
)

type trackBotLeaveCommandHandler struct {
	serverRepository repository.ServerRepository
}

func (h *trackBotLeaveCommandHandler) Handle(s *discordgo.Session, g *discordgo.GuildDelete) {
	_ = h.serverRepository.Delete(entity.CreateServer(g.Guild))
}

func CreateTrackBotLeaveCommand(serverRepository repository.ServerRepository) GuildDeleteHandler {
	return &trackBotLeaveCommandHandler{
		serverRepository: serverRepository,
	}
}
