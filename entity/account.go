package entity

type Account struct {
	Identifier        string
	PlatformAccountId string
	PlatformId        string
	ProfileId         string
}

func CreateAccountFromProfile(profile *Profile, platform *Platform, userAccount string) *Account {
	return &Account{
		Identifier:        profile.Identifier + "_" + platform.Identifier,
		PlatformAccountId: userAccount,
		PlatformId:        platform.Identifier,
		ProfileId:         profile.Identifier,
	}
}
