package models

type User struct {
	BaseModel
	Username string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`

	Addresses []Address `gorm:"foreignKey:UserID"`
}
