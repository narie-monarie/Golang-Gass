package models

type Cat struct {
	Id      int    `json:"cat_id"`
	Name    string `json:"cat_name"`
	CatType string `json:"cat_type"`
}
