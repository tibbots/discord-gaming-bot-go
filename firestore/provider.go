package firestore

import (
	"context"
	"firebase.google.com/go"
	"github.com/tibbots/discord-gaming-bot-go/environment"
	"github.com/tibbots/discord-gaming-bot-go/logging"
	"google.golang.org/api/option"
)

type firestoreClient struct {
	store *firebase.App
}

var firestore Firestore

func init() {
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(environment.GetEnvironment().FirestoreCredentials))
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
func (f *firestoreClient) Users() {
	_, err := f.store.Firestore(context.Background())
	if err != nil {
		logging.Fatal().
			Err(err).
			Msg("unable to connect to firestore.")
	}
	return
}

func GetFirestore() Firestore {
	return firestore
}
