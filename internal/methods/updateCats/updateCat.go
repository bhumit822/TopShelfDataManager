package updatecats

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bhumit/topShelfDataManager/internal/clients"
	// "google.golang.org/api/iterator"
)

func UpdateCateGoriesAndSubCateGories() {
	firestorre, err := clients.App.Firestore(context.Background())
	fmt.Println("1 length of docs is==> ")

	if err == nil {
		iter, err := firestorre.Collection("Products").Documents(context.Background()).GetAll()

		if err != nil {

			fmt.Println("error get docs ==> " + err.Error())

			return
		}

		for i, doc := range iter {

			jsonData, err := json.Marshal(doc.Data())

			if err != nil {
				fmt.Println("get doc data error ==> " + err.Error())
				break
			}
			fmt.Print("doc id===> " + string(jsonData))
			fmt.Print("doc id %d ===> "+string(jsonData), i)

			if i == 10 {
				break
			}

		}

		// var count int
		// for {
		// 	_, err := iter.Next()
		// 	if err == iterator.Done {
		// 		break
		// 	}
		// 	if err != nil {
		// 		fmt.Printf("errrr " + err.Error())
		// 	}

		// 	count++
		// }

		// fmt.Println("2 length of docs is==> %d", count)
	}
	fmt.Println("3 length of docs is==> ")
}
