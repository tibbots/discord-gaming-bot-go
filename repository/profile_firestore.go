package repository

import (
	"context"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
	"github.com/tibbots/discord-gaming-bot-go/logging"
)

type firestoreProfileRepository struct {
	firestore firestore.Firestore
}

func GetFirestoreProfileRepository(firestore firestore.Firestore) ProfileRepository {
	return &firestoreProfileRepository{
		firestore: firestore,
	}
}

func (r *firestoreProfileRepository) Persist(user *entity.Profile) error {
	ctx := context.Background()
	client, err := r.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.Identifier).Set(ctx, user)

	return err
}
func (r *firestoreProfileRepository) GetBy(profile *entity.Profile) (bool, *entity.Profile, error) {
	ctx := context.Background()
	client, err := r.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return false, nil, err
	}
	defer client.Close()

	foundProfileDoc, err := client.Collection("users").Doc(profile.Identifier).Get(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to query profile " + profile.Identifier)
		return false, nil, err
	}

	if !foundProfileDoc.Exists() {
		return false, nil, nil
	}

	foundProfile := &entity.Profile{}
	err = foundProfileDoc.DataTo(foundProfile)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to map found profile " + profile.Identifier)
		return true, nil, err
	}

	return true, foundProfile, nil
}

func (r *firestoreProfileRepository) Delete(profile *entity.Profile) error {
	ctx := context.Background()
	client, err := r.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(profile.Identifier).Delete(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to delete profile " + profile.Identifier)
		return err
	}

	return nil
}
