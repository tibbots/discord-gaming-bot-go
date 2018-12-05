package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/help"
	"github.com/tibbots/discord-gaming-bot-go/repository"
	"strings"
)

type addAccountCommandHandler struct {
	accountRepository  repository.AccountRepository
	profileRepository  repository.ProfileRepository
	platformRepository repository.PlatformRepository
}

func (h *addAccountCommandHandler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := GetMessageCreatedEvent(s, m)
	if !event.shouldBeHandled() || !event.isTalkingToMe() || !event.isCommand("add account") {
		return
	}

	fields := strings.Fields(strings.NewReplacer("add account", "").Replace(event.GetPlainContent()))
	if len(fields) != 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Please provide a platform (eg Steam) and an account-id (eg yourSteamId) for adding an account.")
		return
	}

	platforms, err := h.platformRepository.GetAll()
	if err != nil {
		event.LogError(err, "add account command failed")
		return
	}

	selectedPlatform := strings.ToLower(fields[0])
	selectedAccount := fields[1]
	var matchingPlatform *entity.Platform = nil
	for _, platform := range platforms {
		if strings.EqualFold(platform.Command, selectedPlatform) {
			matchingPlatform = platform
		}
	}

	if matchingPlatform == nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "I am sorry, i dont know anything about a platform called '"+selectedPlatform+"'")
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, help.GetMessages().SupportedPlatforms)
		return
	}

	profile := entity.CreateProfileFromUser(m.Author)
	account := entity.CreateAccountFromProfile(profile, matchingPlatform, selectedAccount)

	err = h.profileRepository.Persist(profile)
	if err != nil {
		event.LogErrorAndRespond(err, "add account command failed")
		return
	}

	_ = h.accountRepository.Persist(account)
	if err != nil {
		event.LogErrorAndRespond(err, "add account command failed")
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, "Yippie, your account has been successfully added! Now you could proceed with the following:")
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, help.GetMessages().NextSteps)
	_, _ = s.ChannelMessageSend(m.ChannelID, help.GetMessages().LinkPreviewHint)
}

func CreateAddAccountCommandHandler(accountRepository repository.AccountRepository,
	profileRepository repository.ProfileRepository, platformRepository repository.PlatformRepository) MessageCreatedHandler {
	return &addAccountCommandHandler{
		accountRepository:  accountRepository,
		profileRepository:  profileRepository,
		platformRepository: platformRepository,
	}
}
