package uploaddata

import (
	"strings"
	"sync"

	firebase "firebase.google.com/go"
)

func getImageListfromString(item string, app *firebase.App) []string {
	var wg sync.WaitGroup

	values := strings.Split(item, ",")
	var items []string
	for value := range values {

		items = append(items, "data/images/"+values[value])

	}

	var urls []string
	for i := range items {
		wg.Add(1)
		urls = append(urls, uploadFileToFireStorage(app, items[i], "products/images/", &wg))

	}

	wg.Wait()
	return urls

}
