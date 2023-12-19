package clients

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
)

type FirestoreClient struct {
	mu     sync.Mutex
	client *firestore.Client
}

func (fc *FirestoreClient) GetClient() (*firestore.Client, error) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	// If the client is not initialized, create a new one
	if fc.client == nil {
		ctx := context.Background()
		client, err := firestore.NewClient(ctx, "topshelf-d392c", ClientOption)
		if err != nil {
			return nil, err
		}
		fc.client = client
	}

	return fc.client, nil
}
