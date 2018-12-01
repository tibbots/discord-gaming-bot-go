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
}

// Environment singleton
var Get *environment

var (
	token       string
	credentials string
)

// inits environment
func init() {
	initParams()

	env := &environment{
		ProjectName:          "discord-gaming-bot",
		ProjectVersion:       "0.2.0",
		ProjectUrl:           "https://github.com/tibbots/discord-gaming-bot-go",
		BotToken:             token,
		FirestoreCredentials: credentials,
	}
	Get = env
}

func initParams() {
	flag.StringVar(&token, "token", "", "bot token")
	flag.StringVar(&credentials, "credentials", "", "firestore credentials")
	flag.Parse()

	if token == "" {
		token = os.Getenv("TOKEN")
	}

	if credentials == "" {
		credentials = os.Getenv("CREDENTIALS")
	}
}
