package help

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/environment"
)

type messages struct {
	SupportedPlatforms *discordgo.MessageEmbed
	NextSteps          *discordgo.MessageEmbed
	Help               *discordgo.MessageEmbed
	Error              *discordgo.MessageEmbed
	LinkPreviewHint    string
}

var messagesInstance *messages

func init() {
	messagesInstance = &messages{
		LinkPreviewHint: "*If you think my previous answer was incomplete, this is most likely because '**Link-Preview**' is disabled in your user-settings. Enable it under **Settings -> Text & Images -> Link Preview***",
		Error: &discordgo.MessageEmbed{
			Title:       "Oops, something went horribly wrong on my side!",
			Description: "Stay tuned, our developers will have a look at it.",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},
		},
		NextSteps: &discordgo.MessageEmbed{
			Title: "Next steps:",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Ask me in any channel to show your personal or a friends Profile (or direct message me)",
					Value: "@MyNameOnYourServer show profile @YourFriend ",
				},
				{
					Name:  "Add a gaming-account to your profile (or direct message me), for example:",
					Value: "@MyNameOnYourServer add account steam your-steam-id",
				},
			},
		},
		SupportedPlatforms: &discordgo.MessageEmbed{
			Title: "Supported Platforms are:",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Steam",
					Value: "add account steam yourSteamId",
				},
				{
					Name:  "Uplay",
					Value: "add account uplay yourUplayId",
				},
				{
					Name:  "Origin",
					Value: "add account origin yourOriginId",
				},
				{
					Name:  "Battlenet",
					Value: "add account battlenet yourBattlenetId",
				},
				{
					Name:  "Microsoft Id",
					Value: "add account xbox yourXboxId",
				},
				{
					Name:  "Playstation Network",
					Value: "add account psn yourPsnId",
				},
			},
		},

		Help: &discordgo.MessageEmbed{
			Title: "Discord Gaming Bot Manual",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Creating a Profile",
					Value: "@MyNameOnYourServer create profile (or via direct message)",
				},
				{
					Name:  "Deleting a Profile",
					Value: "@MyNameOnYourServer delete profile (or via direct message)",
				},
				{
					Name:  "Inspecting your Profile",
					Value: "@MyNameOnYourServer show profile (or via direct message)",
				},
				{
					Name:  "Inspecting another persons Profile",
					Value: "@MyNameOnYourServer show profile @YourFriend (or via direct message)",
				}, {
					Name:  "Adding an platform account (like Steam, Origin or whatsoever)",
					Value: "@MyNameOnYourServer add account [platform] [account-id] (or via direct message)",
				},
			},
		},
	}
}
func GetMessages() *messages {
	return messagesInstance
}
