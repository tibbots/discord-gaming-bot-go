package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"github.com/tibbots/discord-gaming-bot-go/repository"
)

type deleteProfileCommandHandler struct {
	profileRepository repository.ProfileRepository
}

func (h *deleteProfileCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || !event.isCommand("delete profile") {
		return
	}

	err := h.profileRepository.Delete(entity.CreateProfileFromUser(m.Author))
	if err != nil {
		logging.Error().
			Err(err).
			Msg("delete profile command failed")
		_, _ = s.ChannelMessageSend(m.ChannelID, "Oops, something went wrong on my side. Unfortunately i was not able to delete your profile, please try again later.")
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Profile successfully deleted!")
	}
}

func CreateDeleteProfileCommandHandler(profileRepository repository.ProfileRepository) MessageCreatedHandler {
	return &deleteProfileCommandHandler{
		profileRepository: profileRepository,
	}
}
