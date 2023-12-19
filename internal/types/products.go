package types

type Product struct {
	Link          string   `json:"link"`
	MainCategory  string   `json:"main_category"`
	SubCategory   string   `json:"sub_category"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Use           string   `json:"use"`
	Result        string   `json:"result"`
	Content       string   `json:"content"`
	CategoryId    string   `json:"categoryId"`
	SubCategoryId string   `json:"subCategoryId"`
	Price         string   `json:"price"`
	ImagesUrls    []string `json:"images_urls"`
	SearchKey     string   `json:"searchKey"`
	SearchContent []string `json:"searchContent"`
}
