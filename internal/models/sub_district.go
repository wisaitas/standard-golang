package models

type SubDistrict struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	DistrictID int    `json:"district_id"`
	ZipCode    int    `json:"zip_code"`
}
