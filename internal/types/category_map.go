// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    categoryMap, err := UnmarshalCategoryMap(bytes)
//    bytes, err = categoryMap.Marshal()

package types

import "encoding/json"

type CategoryMap []CategoryMapElement

func UnmarshalCategoryMap(data []byte) (CategoryMap, error) {
	var r CategoryMap
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CategoryMap) MarshalCatMap() ([]byte, error) {
	return json.Marshal(r)
}

type CategoryMapElement struct {
	SubCat []SubCat `json:"SubCat,omitempty"`
	Cat    *string  `json:"Cat,omitempty"`
}

type SubCat struct {
	TopShelfSubCategories []string `json:"topShelfSubCategories,omitempty"`
	ProductSubCate        []string `json:"productSubCate,omitempty"`
}
