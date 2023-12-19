package uploaddata

import (
	"context"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
)

func addCategories(client *firestore.Client, wg *sync.WaitGroup, item map[string]interface{}, id string, path string) {
	defer wg.Done()

	fmt.Println("addding new ===>" + id)
	docRef := client.Doc(path + id)
	docRef.Set(context.Background(), item)
	fmt.Printf("category added with doc id ===>  %s*----\n", docRef.ID)
	log.Printf("Goroutine for item %+v is about to exit", item)

}
