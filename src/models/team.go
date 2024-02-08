package models

type Team struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	UUID    string `json:"uuid" gorm:"unique"`
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
}
