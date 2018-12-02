package firestore

import "firebase.google.com/go"

type Firestore interface {
	App() *firebase.App
}
