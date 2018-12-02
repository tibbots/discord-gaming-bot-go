package main

import (
	"github.com/bwmarrin/discordgo"
	env "github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/handler"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	validateParams()

	discord, err := discordgo.New("Bot " + env.GetEnvironment().BotToken)
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to bot. Did you provide a valid token?")
		return
	}

	discord.AddHandler(handler.GetAddAccountHandler().Handle)

	err = discord.Open()
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to bot. Did you provide a valid token?")
		return
	}

	logging.Info().Str("bot", env.GetEnvironment().ProjectName).Str("version", env.GetEnvironment().ProjectVersion).Msg("bot has been successfully started.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()

}

func validateParams() {
	if env.GetEnvironment().BotToken == "" {
		logging.Fatal().Msg("Did you forget to provide a Bot-Token?")
	}

	if env.GetEnvironment().FirestoreCredentials == "" {
		logging.Fatal().Msg("Did you forget to provide the firestore credentials?")
	}
}
