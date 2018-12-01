package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	env "github.com/tibbots/discord-gaming-bot-go/environment"
)

func init() {
	zerolog.TimeFieldFormat = ""
}

func Fatal() *zerolog.Event {
	return apply(log.Fatal())
}

func Info() *zerolog.Event {
	return apply(log.Info())
}

func Debug() *zerolog.Event {
	return apply(log.Debug())
}

func Error() *zerolog.Event {
	return apply(log.Error())
}

func apply(event *zerolog.Event) *zerolog.Event {
	return event.
		Str("bot", env.Get.ProjectName).
		Str("version", env.Get.ProjectVersion)
}
