package repository

import (
	"context"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"google.golang.org/api/iterator"
)

type firestoreAccountProviderRepository struct {
	firestore  firestore.Firestore
	collection string
}

func (r *firestoreAccountProviderRepository) getAll() ([]*entity.AccountProvider, error) {
	ctx := context.Background()
	client, err := r.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return nil, err
	}
	defer client.Close()

	providerDocs := client.Collection(r.collection).Documents(ctx)
	foundProviders := make([]*entity.AccountProvider, 0)
	defer providerDocs.Stop()
	for {
		doc, err := providerDocs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to query account data")
			return nil, err
		}
		foundProvider := &entity.AccountProvider{}
		err = doc.DataTo(foundProvider)
		foundProviders = append(foundProviders, foundProvider)
	}

	return foundProviders, nil
}

func GetFirestoreAccountProviderRepository(firestore firestore.Firestore) AccountProviderRepository {
	return &firestoreAccountProviderRepository{
		firestore:  firestore,
		collection: "platforms",
	}
}
