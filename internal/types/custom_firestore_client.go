package types

import (
	"sync"

	"cloud.google.com/go/firestore"
)

type FirestoreClient struct {
	mu     sync.Mutex
	client *firestore.Client
}
