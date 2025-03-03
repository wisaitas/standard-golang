package models

type District struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}
