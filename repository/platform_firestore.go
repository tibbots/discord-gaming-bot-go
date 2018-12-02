package repository

import (
	"context"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"google.golang.org/api/iterator"
)

type firestorePlatformRepository struct {
	firestore  firestore.Firestore
	collection string
}

func (r *firestorePlatformRepository) GetAll() ([]*entity.Platform, error) {
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
	foundProviders := make([]*entity.Platform, 0)
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
		foundProvider := &entity.Platform{}
		err = doc.DataTo(foundProvider)
		foundProviders = append(foundProviders, foundProvider)
	}

	return foundProviders, nil
}

func GetFirestoreAccountProviderRepository(firestore firestore.Firestore) PlatformRepository {
	return &firestorePlatformRepository{
		firestore:  firestore,
		collection: "platforms",
	}
}
