package handler

import "github.com/bwmarrin/discordgo"

type GuildDeleteHandler interface {
	Handle(s *discordgo.Session, g *discordgo.GuildDelete)
}
