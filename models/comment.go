package models

type Comment struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Text string `json:"comment"`
}
