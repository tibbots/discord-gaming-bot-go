package main

import (
	"github.com/bwmarrin/discordgo"
	env "github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/logging"
)

func main() {
	validateParams()
	logging.Info().Str("bot", env.Get.ProjectName).Str("version", env.Get.ProjectVersion).Msg("bot has been successfully started.")

	discord, err := discordgo.New("Bot " + env.Get.BotToken)
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to bot. Did you provide a valid token?")
		return
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		logging.Info().Msg("message create handler has been invoked")
	})
	err = discord.Open()
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to bot. Did you provide a valid token?")
		return
	}

}

func validateParams() {
	if env.Get.BotToken == "" {
		logging.Fatal().Msg("Did you forget to provide a Bot-Token?")
	}

	if env.Get.FirestoreCredentials == "" {
		logging.Fatal().Msg("Did you forget to provide the firestore credentials?")
	}
}
