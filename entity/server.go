package entity

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Server struct {
	Identifier string
	Name       string
	Region     string
	Members    int
	Channels   int
	Created    int64
	Modified   int64
	Deleted    int64
}

func CreateServer(guild *discordgo.Guild) *Server {
	now := time.Now().Unix()
	return &Server{
		Identifier: guild.ID,
		Name:       guild.Name,
		Created:    now,
		Modified:   now,
		Members:    guild.MemberCount,
		Region:     guild.Region,
		Channels:   len(guild.Channels),
	}
}
