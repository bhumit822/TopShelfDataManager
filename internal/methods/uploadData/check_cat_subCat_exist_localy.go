package uploaddata

import "github.com/bhumit/topShelfDataManager/internal/types"

func isCategoryExist(name string, categories []types.Category) (bool, *types.Category) {
	for _, p := range categories {
		if p.Name == name {
			return true, &p
		}
	}
	return false, nil
}
func isSubCategoryExist(name string, categories []types.SubCategory) (bool, *types.SubCategory) {
	for _, p := range categories {
		if p.Name == name {
			return true, &p
		}
	}
	return false, nil
}
