package models

type Province struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}
