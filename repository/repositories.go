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

type PlatformRepository interface {
	GetAll() ([]*entity.Platform, error)
}

type ProfileRepository interface {
	Persist(user *entity.Profile) error
	GetBy(user *entity.Profile) (bool, *entity.Profile, error)
	Delete(user *entity.Profile) error
}

type ServerRepository interface {
	Persist(server *entity.Server) error
	Delete(server *entity.Server) error
	GetAll() ([]*entity.Server, error)
}

func GetAccountRepository() AccountRepository {
	return GetFirestoreAccountRepository(firestore.GetFirestore())
}

func GetProfileRepository() ProfileRepository {
	return GetFirestoreProfileRepository(firestore.GetFirestore())
}

func GetPlatformRepository() PlatformRepository {
	return GetFirestoreAccountProviderRepository(firestore.GetFirestore())
}

func GetServerRepository() ServerRepository {
	return GetFirestoreServerRepository(firestore.GetFirestore())
}
