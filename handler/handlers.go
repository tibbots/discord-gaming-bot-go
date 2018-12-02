package handler

import "github.com/bwmarrin/discordgo"

type MessageCreatedHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
}
