package updatecats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/bhumit/topShelfDataManager/internal/clients"
	"github.com/bhumit/topShelfDataManager/internal/types"
	// "github.com/bhumit/topShelfDataManager/internal/types"
	// "google.golang.org/api/iterator"
)

type ProductCategory struct {
	Cat    string        `json:"Cat"`
	SubCat []SubCategory `json:"SubCat"`
}

type SubCategory struct {
	TopShelfSubCategories string   `json:"topShelfSubCategories"`
	ProductSubCate        []string `json:"productSubCate"`
}

func UpdateCateGoriesAndSubCateGories() {
	// firestoreClient := &clients.FirestoreClient{}
	// cat, subCat, err := getCategoryAndSubCategory("Gesichtspflege")

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return
	// }
	// fmt.Printf("category===> " + *cat)
	// fmt.Printf("\n\nSubcategory===> " + *subCat)

	// return
	firestorre, err := clients.App.Firestore(context.Background())
	var wg sync.WaitGroup

	if err == nil {
		iter, err := firestorre.Collection("Products").Documents(context.Background()).GetAll()

		if err != nil {

			fmt.Println("error get docs ==> " + err.Error())

			return
		}

		for i, doc := range iter {
			wg.Add(1)
			jsonData, err := json.Marshal(doc.Data())

			if err != nil {
				fmt.Println("get doc data error ==> " + err.Error())
				break
			}
			// fmt.Print("doc id===> " + string(jsonData))
			// fmt.Print("doc id %d ===> "+string(jsonData), i)
			var product types.Product
			json.Unmarshal(jsonData, &product)

			go func(doc *firestore.DocumentSnapshot) {
				wg.Done()
				// client, err := firestoreClient.GetClient()
				// if err != nil {
				// 	log.Printf("Error getting Firestore client: %v", err)

				// 	wg.Done()
				// 	return
				// } else {
				// getCategoryAndSubCategory(product.MainCategory, product.SubCategory)
				cat, subCat, err := getCategoryAndSubCategory(product.MainCategory, product.SubCategory)

				if err != nil {
					fmt.Print("cat sub cate" + err.Error())
					return
				}
				// fmt.Printf("category===> " + *cat)
				// fmt.Printf("\n\nSubcategory===> " + *subCat)

				updateData := map[string]interface{}{
					"topshelfCategoryId":    cat,
					"topshelfSubCategoryId": subCat,
					"updatedOn":             time.Now(),
				}

				_, err = doc.Ref.Set(context.Background(), updateData, firestore.MergeAll)

				if err != nil {
					log.Fatalf("Failed to update document: %v", err)
				}
				docLen := len(iter)
				fmt.Printf("\nDocument updated successfully! %d -outOf- %d", i, docLen)
				fmt.Print("\nDocument updated successfully!   ==>" + doc.Ref.ID + "\n\n")

				// }

			}(doc)

			// if i == 10 {
			// 	break
			// }
			if i%100 == 0 {
				wg.Wait()
			}

		}

	}
	wg.Wait()
	fmt.Println("3 length of docs is==> ")
}

func getCategoryAndSubCategory(Mcat string, subCat string) (*string, *string, error) {
	var cat *string = nil
	var subCatNew *string = nil
	catData, errd := getCatData()

	if errd != nil {

		fmt.Print("inside cat sub cat")
		return nil, nil, errd

	}

	for _, cat_ := range catData {

		for _, subCat_ := range cat_.SubCat {

			for _, pSubCat_ := range subCat_.ProductSubCate {

				if strings.ToLower(pSubCat_) == strings.ToLower(subCat) {

					_subCat := strings.ToUpper(subCat_.TopShelfSubCategories)
					_cat := strings.ToUpper(cat_.Cat)
					subCatNew = &_subCat
					cat = &_cat
					break

				}
			}
			if cat != nil && subCatNew != nil {
				break
			}

		}
		if cat != nil && subCatNew != nil {
			break
		}

	}
	if cat == nil && subCatNew == nil {
		return nil, nil, errors.New("no cat and sub cat found.==>" + subCat + "====" + Mcat + "\n")
	}
	return cat, subCatNew, nil

}

func getCatData() ([]ProductCategory, error) {
	var jsonData = []byte(`
	[
		{
			"SubCat": [
				{"topShelfSubCategories": "TEINT", "productSubCate": ["Teint"]},
				{"topShelfSubCategories": "LIPPEN", "productSubCate": ["Lippen"]},
				{"topShelfSubCategories": "PINSEL", "productSubCate": ["Pinsel und Schwamm", "Pinsel"]},
				{"topShelfSubCategories": "AUGEN", "productSubCate": ["Augen"]},
				{"topShelfSubCategories": "AUGENBRAUEN", "productSubCate": ["Augenbrauen"]},
				{"topShelfSubCategories": "NÄGEL", "productSubCate": ["Nägel"]},
				{"topShelfSubCategories": "MAKE-UP ACCESSOIRES", "productSubCate": ["Accessoires Make Up", "Zubehör"]},
				{"topShelfSubCategories": "SETS", "productSubCate": ["Make-up Sets"]}
			],
			"Cat": "MAKE UP"
		},
		{
			"SubCat": [
				{"topShelfSubCategories": "GESICHTSPFLEGE", "productSubCate": ["Gesichtspflege"]},
				{"topShelfSubCategories": "Gesichtsreinigung", "productSubCate": ["Gesichtsserum", "Gesichtsreinigung"]},
				{"topShelfSubCategories": "MASKEN", "productSubCate": ["Masken", "Gesichtsmasken"]},
				{"topShelfSubCategories": "SONNE & SCHUTZ", "productSubCate": ["Sonnen + Sonnenschutz", "Sonne & Schutz"]},
				{"topShelfSubCategories": "Hauttypen", "productSubCate": ["Pflege nach Hautbedürfnis", "Hauttypen"]},
				{"topShelfSubCategories": "BEAUTY TOOLS", "productSubCate": ["Accessoires", "Beauty Tools & Zubehör"]},
				{"topShelfSubCategories": "AUGENPFLEGE", "productSubCate": ["Augenpflege"]},
				{"topShelfSubCategories": "LIPPENPFLEGE", "productSubCate": ["Lippenpflege"]}
			],
			"Cat": "GESICHT"
		},
		{
			"SubCat": [
				{"topShelfSubCategories": "DAMENDÜFTE", "productSubCate": ["Damendüfte"]},
				{"topShelfSubCategories": "HERRENDÜFTE", "productSubCate": ["Herrendüfte"]},
				{"topShelfSubCategories": "UNISEX DÜFTE", "productSubCate": ["Unisex Düfte"]},
				{"topShelfSubCategories": "NISCHENDÜFTE", "productSubCate": ["Nischenparfüm", "Nischendüfte"]},
				{"topShelfSubCategories": "PARFUM SETS", "productSubCate": ["Parfum Sets"]}
			],
			"Cat": "PARFUM"
		},
		{
			"SubCat": [
				{"topShelfSubCategories": "SHAMPOO", "productSubCate": []},
				{"topShelfSubCategories": "TROCKENSHAMPOO", "productSubCate": []},
				{"topShelfSubCategories": "CONDITIONER", "productSubCate": []},
				{"topShelfSubCategories": "LEAVE-IN-BEHANDLUNG", "productSubCate": []},
				{"topShelfSubCategories": "HAARKUR & -MASKE", "productSubCate": []},
				{"topShelfSubCategories": "HAAR-ÖLE & SEEREN", "productSubCate": []},
				{"topShelfSubCategories": "HAARPARFÜM", "productSubCate": []},
				{"topShelfSubCategories": "KOPFHAUTPFLEGE", "productSubCate": []},
				{"topShelfSubCategories": "SONNENSCHUTZ", "productSubCate": []},
				{"topShelfSubCategories": "HAARPFLEGESETS", "productSubCate": ["Haarpflege"]}
			],
			"Cat": "HAARPFLEGE"
		},
		{
			"SubCat": [
				{"topShelfSubCategories": "HAARSPRAY", "productSubCate": []},
				{"topShelfSubCategories": "HAARWACHS & CREME", "productSubCate": []},
				{"topShelfSubCategories": "HAARGEL", "productSubCate": []},
				{"topShelfSubCategories": "HAARMOUSSE", "productSubCate": []},
				{"topShelfSubCategories": "HITZESCHUTZ", "productSubCate": []},
				{"topShelfSubCategories": "STYLINGSPRAYS", "productSubCate": []}
			],
			"Cat": "HAARSTYLING"
		}
	]
`)

	var productCategories []ProductCategory

	err := json.Unmarshal(jsonData, &productCategories)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, errors.New("unable to get a valid value")

	}

	return productCategories, nil
}
