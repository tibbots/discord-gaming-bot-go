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
	Created    int64
	Modified   int64
	Deleted    int64
}

func CreateServer(guild *discordgo.Guild) *Server {
	return &Server{
		Identifier: guild.ID,
		Name:       guild.Name,
		Created:    time.Now().Unix(),
		Modified:   time.Now().Unix(),
		Members:    guild.MemberCount,
		Region:     guild.Region,
	}
}
