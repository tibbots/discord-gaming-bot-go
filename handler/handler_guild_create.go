package handler

import "github.com/bwmarrin/discordgo"

type GuildCreateHandler interface {
	Handle(s *discordgo.Session, g *discordgo.GuildCreate)
}
