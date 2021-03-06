package main

import (
	"github.com/bwmarrin/discordgo"
	env "github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/handler"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"github.com/tibbots/discord-gaming-bot-go/repository"
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

	createProfileCommand := handler.CreateCreateProfileCommandHandler(repository.GetProfileRepository())
	deleteProfileCommand := handler.CreateDeleteProfileCommandHandler(repository.GetProfileRepository())
	showProfileCommand := handler.CreateShowProfileCommandHandler(repository.GetAccountRepository(), repository.GetPlatformRepository(), repository.GetProfileRepository())
	addAccountCommand := handler.CreateAddAccountCommandHandler(
		repository.GetAccountRepository(),
		repository.GetProfileRepository(),
		repository.GetPlatformRepository())
	helpCommand := handler.CreateHelpCommandHandler()
	versionCommand := handler.CreateVersionCommandHandler(repository.GetServerRepository())
	trackBotJoins := handler.CreateTrackBotJoinsCommand(repository.GetServerRepository())
	trackBotLeaves := handler.CreateTrackBotLeaveCommand(repository.GetServerRepository())
	pushStats := handler.CreatePushStatsCommandHandler(repository.GetServerRepository())

	discord.AddHandler(addAccountCommand.Handle)
	discord.AddHandler(createProfileCommand.Handle)
	discord.AddHandler(deleteProfileCommand.Handle)
	discord.AddHandler(showProfileCommand.Handle)
	discord.AddHandler(helpCommand.Handle)
	discord.AddHandler(versionCommand.Handle)
	discord.AddHandler(trackBotJoins.Handle)
	discord.AddHandler(trackBotLeaves.Handle)
	discord.AddHandler(pushStats.Handle)

	err = discord.Open()
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to bot. Did you provide a valid token?")
		return
	}

	logging.Info().Str("bot", env.GetEnvironment().ProjectName).Str("version", env.GetEnvironment().ProjectVersion).Msg("bot is up and running.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logging.Info().Str("bot", env.GetEnvironment().ProjectName).Str("version", env.GetEnvironment().ProjectVersion).Msg("bot is shutting down.")
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
