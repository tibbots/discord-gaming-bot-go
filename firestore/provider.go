package firestore

import (
	"context"
	"firebase.google.com/go"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"google.golang.org/api/option"
	"path/filepath"
)

type firestoreClient struct {
	store *firebase.App
}

var firestore Firestore

func init() {
	credentialsFile, err := filepath.Abs(environment.GetEnvironment().FirestoreCredentials)
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to firestore. Did you provide a valid path to JSON credentials?")
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to firestore. Did you provide a valid path to JSON credentials?")
	}
	firestore = &firestoreClient{
		store: app,
	}
}

func (f *firestoreClient) App() *firebase.App {
	return f.store
}

func GetFirestore() Firestore {
	return firestore
}
