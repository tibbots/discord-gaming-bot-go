package repository

import (
	"context"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"google.golang.org/api/iterator"
)

type firestoreAccountRepository struct {
	firestore firestore.Firestore
}

func GetFirestoreAccountRepository(firestore firestore.Firestore) AccountRepository {
	return &firestoreAccountRepository{
		firestore: firestore,
	}
}

func (a *firestoreAccountRepository) Persist(account entity.Account) error {
	ctx := context.Background()
	client, err := a.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	_, err = client.Collection("accounts").Doc(account.Identifier).Set(ctx, account)

	return err
}

func (a *firestoreAccountRepository) GetByProfile(user entity.Profile) ([]*entity.Account, error) {
	ctx := context.Background()
	client, err := a.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return nil, err
	}
	defer client.Close()

	existingAccounts := client.Collection("accounts").Where("userId", "==", user.Identifier).Documents(ctx)
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

func (a *firestoreAccountRepository) Delete(user entity.Profile) error {
	ctx := context.Background()
	client, err := a.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	foundAccounts, err := a.GetByProfile(user)
	if err != nil {
		return err
	}

	for _, foundAccount := range foundAccounts {
		_, err := client.Collection("accounts").Doc(foundAccount.Identifier).Delete(ctx)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to delete account data for user " + user.Identifier)
			return err
		}
	}

	return nil
}
