package main

import (
	"github.com/bhumit/topShelfDataManager/internal/clients"
	updatecats "github.com/bhumit/topShelfDataManager/internal/methods/updateCats"
	// uploaddata "github.com/bhumit/topShelfDataManager/internal/methods/uploadData"
)

func main() {
	clients.InitFirebase()
	updatecats.UpdateCateGoriesAndSubCateGories()
	// uploaddata.UploadData()

}
