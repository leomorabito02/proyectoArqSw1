package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Nombre   string `gorm:"type:varchar(300);not null"`
	Apellido string `gorm:"type:varchar(300);not null"`
	Email    string `gorm:"type:varchar(500);not null;unique"`
	Password string `gorm:"type:varchar(200);not null"`
	Tipo     bool   `gorm:"type:boolean;not null"`
	Dni      int    `gorm:"type int;not null"`
}

type Users []User
