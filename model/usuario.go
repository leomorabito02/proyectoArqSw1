package model

type Usuario struct {
	Id       int    `gorm:"primaryKey"`
	Nombre   string `gorm:"type:varchar(20);not null"`
	Apellido string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"type:varchar(16);not null"`
	Tipo     bool   `gorm:"type:boolean();not null"`
	Dni      int    `gorm:"type int(10);unsigned;not null"`
}

type Usuarios []Usuario
