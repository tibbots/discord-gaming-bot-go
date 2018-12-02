package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"github.com/tibbots/discord-gaming-bot-go/repository"
	"strings"
)

type addAccountCommandHandler struct {
	accountRepository  repository.AccountRepository
	profileRepository  repository.ProfileRepository
	platformRepository repository.PlatformRepository
	failureMessage     *discordgo.MessageEmbed
	successMessage     *discordgo.MessageEmbed
	supportedPlatforms *discordgo.MessageEmbed
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
		logging.Error().
			Err(err).
			Msg("add account command failed")
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
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.supportedPlatforms)
		return
	}

	profile := entity.CreateProfileFromUser(m.Author)
	account := entity.CreateAccountFromProfile(profile, matchingPlatform, selectedAccount)

	err = h.profileRepository.Persist(profile)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("add account command failed")
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.failureMessage)
		return
	}

	_ = h.accountRepository.Persist(account)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("add account command failed")
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.failureMessage)
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, h.successMessage)
}

func CreateAddAccountCommandHandler(accountRepository repository.AccountRepository,
	profileRepository repository.ProfileRepository, platformRepository repository.PlatformRepository) MessageCreatedHandler {
	return &addAccountCommandHandler{
		accountRepository:  accountRepository,
		profileRepository:  profileRepository,
		platformRepository: platformRepository,
		failureMessage: &discordgo.MessageEmbed{
			Title:       "Oops, something went horribly wrong on my side!",
			Description: "Stay tuned, our developers will have a look at it.",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},
		},

		successMessage: &discordgo.MessageEmbed{
			Title: "Your Account has been successfully added. Next steps to take:",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Ask me in any channel to show your personal or a friends Profile (or direct message me)",
					Value: "@MyNameOnYourServer show profile @YourFriend ",
				},
			},
		},

		supportedPlatforms: &discordgo.MessageEmbed{
			Title: "Supported Platforms are:",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "reach us at " + environment.GetEnvironment().ProjectUrl,
			},

			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Steam",
					Value: "add account steam yourSteamId",
				},
				{
					Name:  "Uplay",
					Value: "add account uplay yourUplayId",
				},
				{
					Name:  "Origin",
					Value: "add account origin yourOriginId",
				},
				{
					Name:  "Battlenet",
					Value: "add account battlenet yourBattlenetId",
				},
				{
					Name:  "Microsoft Id",
					Value: "add account xbox yourXboxId",
				},
				{
					Name:  "Playstation Network",
					Value: "add account psn yourPsnId",
				},
			},
		},
	}
}
