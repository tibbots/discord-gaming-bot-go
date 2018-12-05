package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/help"
	"github.com/tibbots/discord-gaming-bot-go/repository"
)

type showProfileCommandHandler struct {
	accountRepository  repository.AccountRepository
	platformRepository repository.PlatformRepository
	profileRepository  repository.ProfileRepository
}

func (h *showProfileCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || (!event.isCommand("show profile") && !event.isCommand("inspect profile")) {
		return
	}

	profile := entity.CreateProfileFromUser(event.GetTargetedUser())

	exists, _, _ := h.profileRepository.GetBy(profile)
	if !exists {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Looks like <@%s> has not yet registered a Profile", profile.DiscordUserId))
		return
	}

	accounts, err := h.accountRepository.GetByProfile(profile)
	if err != nil {
		event.LogErrorAndRespond(err, "show profile command failed")
		return
	}
	platforms, err := h.platformRepository.GetAll()
	if err != nil {
		event.LogErrorAndRespond(err, "show profile command failed")
		return
	}

	platformIdToPlatform := make(map[string]*entity.Platform)
	for _, platform := range platforms {
		platformIdToPlatform[platform.Identifier] = platform
	}

	accountFields := make([]*discordgo.MessageEmbedField, 0)
	for _, account := range accounts {
		accountFields = append(accountFields, &discordgo.MessageEmbedField{
			Name:  platformIdToPlatform[account.PlatformId].Name + fmt.Sprintf(" (%s)", platformIdToPlatform[account.PlatformId].Command),
			Value: account.PlatformAccountId,
		})
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("Gaming Profile of <@%s>", profile.DiscordUserId),
		Fields:      accountFields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
		},
	})

	_, _ = s.ChannelMessageSend(m.ChannelID, help.GetMessages().LinkPreviewHint)

	if err != nil {
		event.LogError(err, "show profile command failed")
	}
}

func CreateShowProfileCommandHandler(accountRepository repository.AccountRepository,
	platformRepository repository.PlatformRepository, profileRepository repository.ProfileRepository) MessageCreatedHandler {
	return &showProfileCommandHandler{
		accountRepository:  accountRepository,
		platformRepository: platformRepository,
		profileRepository:  profileRepository,
	}
}
