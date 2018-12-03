package repository

import (
	"context"
	"github.com/tibbots/discord-gaming-bot-go/entity"
	"github.com/tibbots/discord-gaming-bot-go/firestore"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"time"
)

type firestoreServerRepository struct {
	firestore  firestore.Firestore
	collection string
}

func (r *firestoreServerRepository) Persist(server *entity.Server) error {
	ctx := context.Background()
	client, err := r.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	serverDoc, err := client.Collection(r.collection).Doc(server.Identifier).Get(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("error retrieving server data")
		// dont return here, as err is of type NotFound, although serverDoc is non-nil
	}
	if !serverDoc.Exists() {
		_, err = client.Collection(r.collection).Doc(server.Identifier).Create(ctx, server)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to create server data")
			return err
		}
	} else {
		existingServer := &entity.Server{}
		err = serverDoc.DataTo(existingServer)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to map server data")
			return err
		}

		existingServer.Region = server.Region
		existingServer.Members = server.Members
		existingServer.Name = server.Name
		existingServer.Deleted = 0
		existingServer.Modified = time.Now().Unix()
		_, err = client.Collection(r.collection).Doc(server.Identifier).Set(ctx, existingServer)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to update server data")
			return err
		}
	}

	return nil
}

func (r *firestoreServerRepository) Delete(server *entity.Server) error {
	ctx := context.Background()
	client, err := r.firestore.App().Firestore(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to connect to firestore.")
		return err
	}
	defer client.Close()

	serverDoc, err := client.Collection(r.collection).Doc(server.Identifier).Get(ctx)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("error retrieving server data")
		// dont return here, as err is of type NotFound, although serverDoc is non-nil
	}
	existingServer := server
	if serverDoc.Exists() {
		existingServer = &entity.Server{}
		err = serverDoc.DataTo(existingServer)
		if err != nil {
			logging.Error().
				Err(err).
				Msg("unable to map server data")
			return err
		}

		existingServer.Region = server.Region
		existingServer.Members = server.Members
		existingServer.Name = server.Name
	}

	existingServer.Deleted = time.Now().Unix()
	existingServer.Modified = existingServer.Deleted

	_, err = client.Collection(r.collection).Doc(server.Identifier).Set(ctx, existingServer)
	if err != nil {
		logging.Error().
			Err(err).
			Msg("unable to update server data")
		return err
	}

	return nil
}

func GetFirestoreServerRepository(firestore firestore.Firestore) ServerRepository {
	return &firestoreServerRepository{
		firestore:  firestore,
		collection: "servers",
	}
}
