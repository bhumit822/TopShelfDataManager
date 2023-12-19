package uploaddata

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bhumit/topShelfDataManager/internal/clients"
	"github.com/bhumit/topShelfDataManager/internal/types"
	"github.com/google/uuid"
)

func UploadData() {

	file, err := os.Open("data/data.csv") // Replace with your CSV file path
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var cats []types.Category

	var subCats []types.SubCategory
	// var data []Product

	var wg sync.WaitGroup
	firestoreClient := &clients.FirestoreClient{}

	///===> fetch all categories and sub-categories from product list
	for _, record := range records[1:] {
		isEx, cat := isCategoryExist(record[2], cats)
		if !isEx {

			cat = &types.Category{Id: uuid.NewString(), Name: record[2]}
			fmt.Println("cat len ===>" + cat.Id)
			cats = append(cats, *cat)

		}

		issubEx, subCat := isSubCategoryExist(record[3], subCats)
		if !issubEx {
			subCat = &types.SubCategory{Id: uuid.NewString(), Name: record[3], MainCategoryId: cat.Id}
			subCats = append(subCats, *subCat)
		}
	}

	///=====> add categories to firestore
	for _, catt := range cats {

		fmt.Println(catt.Name + "\n")
		wg.Add(1)
		go func(cattt types.Category) {
			client, err := firestoreClient.GetClient()
			if err != nil {
				log.Printf("Error getting Firestore client: %v", err)
				wg.Done()
				return
			}

			categoriesJson, err := json.Marshal(cattt)
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println("Error:", categoriesJson)

			}

			var catinInterface map[string]interface{}
			json.Unmarshal(categoriesJson, &catinInterface)
			// Print the JSON data

			addCategories(client, &wg, catinInterface, cattt.Id, "category/")

		}(catt)

	}

	///=====> add sub-categories to firestore
	for _, sCatt := range subCats {
		fmt.Println(sCatt.Name + "\n")
		wg.Add(1)
		go func(ssCat types.SubCategory) {
			client, err := firestoreClient.GetClient()
			if err != nil {
				log.Printf("Error getting Firestore client: %v", err)
				wg.Done()
				return
			}

			categoriesJson, err := json.Marshal(ssCat)
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println("Error:", categoriesJson)

			}

			var subCatinInterface map[string]interface{}
			json.Unmarshal(categoriesJson, &subCatinInterface)
			// Print the JSON data

			addCategories(client, &wg, subCatinInterface, ssCat.Id, "subCategory/")

		}(sCatt)

	}
	wg.Wait()

	var wgProduct sync.WaitGroup
	///=======> add products to firestore
	for i, record := range records[1:] {
		wgProduct.Add(2)
		_, cat := isCategoryExist(record[2], cats)
		_, subCat := isSubCategoryExist(record[3], subCats)
		go func(record []string) {
			defer wgProduct.Done()
			client, err := firestoreClient.GetClient()
			if err != nil {
				log.Printf("Error getting Firestore client: %v", err)
				wgProduct.Done()
				return
			}
			addProduct(clients.App, client, &wgProduct, record, cat.Id, subCat.Id)

		}(record)

		if i%100 == 0 {
			wgProduct.Wait()
		}

	}

	wgProduct.Wait()

	fmt.Println("\n\n================  Product uploaded =================")

}

///=================>goroutine for upload subcategories

// // remove element from list of product
// func removeElement(slice []Product, index int) []Product {
// 	return append(slice[:index], slice[index+1:]...)
// }

// func getImagesPaths(path string) []string {
// 	images, err := os.Open(path)

// 	if err != nil {

// 		fmt.Println(err.Error())

// 	}

// 	var fileNames []string

// 	files, err := images.Readdir(-1)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	for _, file := range files {
// 		if file.Mode().IsRegular() {
// 			fileNames = append(fileNames, file.Name())
// 			fmt.Println(file.Name())
// 		}
// 	}

// 	images.Close()

// 	return fileNames
// }

// func getUniqueElements(input []string) []string {
// 	uniqueMap := make(map[string]bool)
// 	uniqueElements := []string{}

// 	for _, element := range input {
// 		if !uniqueMap[element] {
// 			uniqueMap[element] = true
// 			uniqueElements = append(uniqueElements, element)
// 		}
// 	}

// 	return uniqueElements
// }
