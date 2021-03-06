package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
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
		event.LogErrorAndRespond(err, "delete profile command failed")
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, "No problem, your profile has been successfully deleted!")
}

func CreateDeleteProfileCommandHandler(profileRepository repository.ProfileRepository) MessageCreatedHandler {
	return &deleteProfileCommandHandler{
		profileRepository: profileRepository,
	}
}
