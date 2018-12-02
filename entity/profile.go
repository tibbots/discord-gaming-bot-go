package entity

import "github.com/bwmarrin/discordgo"

type Profile struct {
	Identifier      string
	DiscordUserId   string
	DiscordUserName string
}

func CreateProfileFromUser(user *discordgo.User) *Profile {
	return &Profile{
		Identifier:      user.ID,
		DiscordUserId:   user.ID,
		DiscordUserName: user.Username + "#" + user.Discriminator,
	}
}
