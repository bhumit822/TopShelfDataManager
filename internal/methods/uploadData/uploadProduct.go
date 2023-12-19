package uploaddata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/bhumit/topShelfDataManager/internal/types"
)

func addProduct(app *firebase.App, client *firestore.Client, wg *sync.WaitGroup, record []string, catid string, subCatId string) {
	defer wg.Done()
	images := getImageListfromString(record[10], app)

	searchContent := strings.Split(strings.ToLower(record[5]), " ")

	searchContent = append(searchContent, strings.Split(strings.ToLower(record[4]), " ")...)

	product := types.Product{Link: record[1], MainCategory: record[2], CategoryId: catid, SubCategoryId: subCatId, SubCategory: record[3], Name: record[4], SearchKey: strings.ToLower(record[4]), Description: record[5], Use: record[6], Result: record[7], Content: record[8], Price: record[9], ImagesUrls: images, SearchContent: searchContent}

	productJsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Error:", productJsonData)
		return
	}

	var inInterface map[string]interface{}
	json.Unmarshal(productJsonData, &inInterface)

	docRef, _, err := client.Collection("Products").Add(context.Background(), inInterface)
	if err != nil {
		log.Fatalf("error adding document: %v", err)
	}
	fmt.Printf("category added with doc id ===>  %s*----\n", docRef.ID)

}
