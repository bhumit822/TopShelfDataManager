package uploaddata

import (
	"context"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
)

func uploadFileToFireStorage(app *firebase.App, filePath string, destination string, wg *sync.WaitGroup) string {
	defer wg.Done()
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf("error initializing storage client: %v", err)

	}

	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Fatalf("error opening file %s: %v", filePath, err)

	// }
	// defer file.Close()

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalf("error getting default bucket: %v", err)

	}
	data, err := os.ReadFile(filePath)

	var fileName string = strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1]
	wc := bucket.Object(destination + fileName).NewWriter(context.Background())
	if _, err := wc.Write(data); err != nil {
		log.Fatalf("error uploading file: %v", err)

	}
	if err := wc.Close(); err != nil {
		log.Fatalf("error closing writer: %v", err)

	}
	downloadURL, err := bucket.SignedURL(destination+fileName, &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})
	if err != nil {
		log.Fatalf("error getting download URL: %v\n", err)
	}
	// fmt.Printf("File uploaded from %s.\n", filePath)
	// fmt.Printf("File uploaded from %s.\n", filePath)
	// fmt.Printf("download URL ==========> %s.\n", downloadURL)
	return downloadURL
}
