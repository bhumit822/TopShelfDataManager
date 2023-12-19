package types

type Category struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
type SubCategory struct {
	Id             string `json:"id"`
	MainCategoryId string `json:"mainCategoryId"`
	Name           string `json:"name"`
	Image          string `json:"image"`
}
