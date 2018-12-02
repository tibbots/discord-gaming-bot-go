package repository

import (
	"context"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"google.golang.org/api/iterator"
)

type firestoreAccountRepository struct {
	firestore  firestore.Firestore
	collection string
}

func GetFirestoreAccountRepository(firestore firestore.Firestore) AccountRepository {
	return &firestoreAccountRepository{
		firestore:  firestore,
		collection: "platform_accounts",
	}
}

func (a *firestoreAccountRepository) Persist(account *entity.Account) error {
	ctx := context.Background()
	client, err := a.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	_, err = client.Collection(a.collection).Doc(account.Identifier).Set(ctx, account)

	return err
}

func (a *firestoreAccountRepository) GetByProfile(profile *entity.Profile) ([]*entity.Account, error) {
	ctx := context.Background()
	client, err := a.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return nil, err
	}
	defer client.Close()

	existingAccounts := client.Collection(a.collection).Where("ProfileId", "==", profile.Identifier).Documents(ctx)
	foundAccounts := make([]*entity.Account, 0)
	defer existingAccounts.Stop()
	for {
		doc, err := existingAccounts.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to query account data")
			return nil, err
		}
		foundAccount := &entity.Account{}

		err = doc.DataTo(foundAccount)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to map queried account data")
			return nil, err
		}

		foundAccounts = append(foundAccounts, foundAccount)

	}
	return foundAccounts, nil
}

func (a *firestoreAccountRepository) Delete(profile *entity.Profile) error {
	ctx := context.Background()
	client, err := a.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	foundAccounts, err := a.GetByProfile(profile)
	if err != nil {
		return err
	}

	for _, foundAccount := range foundAccounts {
		_, err := client.Collection(a.collection).Doc(foundAccount.Identifier).Delete(ctx)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to delete account data for user " + profile.Identifier)
			return err
		}
	}

	return nil
}
