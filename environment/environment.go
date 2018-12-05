package environment

import (
	"flag"
	"os"
)

// Environment struct
type environment struct {
	ProjectName          string
	ProjectVersion       string
	ProjectUrl           string
	BotToken             string
	FirestoreCredentials string
	DiscordBotsToken     string
}

// Environment singleton
var environmentInstance *environment

var (
	token            string
	credentials      string
	tokenDiscordBots string
)

func GetEnvironment() *environment {
	return environmentInstance
}

// inits environment
func init() {
	initParams()

	env := &environment{
		ProjectName:          "discord-gaming-bot",
		ProjectVersion:       "0.1.2",
		ProjectUrl:           "https://github.com/tibbots/discord-gaming-bot-go",
		BotToken:             token,
		FirestoreCredentials: credentials,
		DiscordBotsToken:     tokenDiscordBots,
	}
	environmentInstance = env
}

func initParams() {
	flag.StringVar(&token, "token", "", "bot token")
	flag.StringVar(&credentials, "credentials", "", "firestore credentials")
	flag.StringVar(&tokenDiscordBots, "tokenDiscordBots", "", "discord-bots token")
	flag.Parse()

	if token == "" {
		token = os.Getenv("TOKEN")
	}

	if credentials == "" {
		credentials = os.Getenv("CREDENTIALS")
	}

	if tokenDiscordBots == "" {
		tokenDiscordBots = os.Getenv("TOKEN_DISCORD_BOTS")
	}
}
