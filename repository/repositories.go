package repository

import (
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
)

type AccountRepository interface {
	Persist(account *entity.Account) error
	GetByProfile(user *entity.Profile) ([]*entity.Account, error)
	Delete(user *entity.Profile) error
}

type AccountProviderRepository interface {
	getAll() ([]*entity.AccountProvider, error)
}

type ProfileRepository interface {
	Persist(user *entity.Profile) error
	GetBy(user *entity.Profile) (bool, *entity.Profile, error)
	Delete(user *entity.Profile) error
}

func GetAccountRepository() AccountRepository {
	return GetFirestoreAccountRepository(firestore.GetFirestore())
}

func GetProfileRepository() ProfileRepository {
	return GetFirestoreProfileRepository(firestore.GetFirestore())
}

func GetAccountProviderRepository() AccountProviderRepository {
	return GetFirestoreAccountProviderRepository(firestore.GetFirestore())
}
