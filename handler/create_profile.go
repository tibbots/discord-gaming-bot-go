package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/repository"
)

type createProfileCommandHandler struct {
	profileRepository repository.ProfileRepository
}

func (h *createProfileCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || !event.isCommand("create profile") {
		return
	}

	err := h.profileRepository.Persist(entity.CreateProfileFromUser(m.Author))
	if err != nil {
		event.LogErrorAndRespond(err, "create profile command failed")
		return
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Profile successfully created!")
	}
}

func CreateCreateProfileCommandHandler(profileRepository repository.ProfileRepository) MessageCreatedHandler {
	return &createProfileCommandHandler{
		profileRepository: profileRepository,
	}
}
