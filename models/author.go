package models

type Author struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
