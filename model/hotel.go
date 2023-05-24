package model

type Hotel struct {
	Id                int     `gorm:"primaryKey"`
	Nombre            string  `gorm:"type:varchar(300);not null"`
	Descripcion       string  `gorm:"type:text(500);not null"`
	Cant_habitaciones int     `gorm:"type:int(400);not null"`
	Valoracion        int     `gorm:"type:int(100);unsigned;not null"`
	Gym               bool    `gorm:"type:boolean();not null"`
	Wifi              bool    `gorm:"type:boolean();not null"`
	Estacionamiento   bool    `gorm:"type:boolean();not null"`
	Bide              bool    `gorm:"type:boolean();not null"`
	Pileta            bool    `gorm:"type:boolean();not null"`
	Precio            float32 `gorm:"type:decimal(400);unsigned;not null"`
}

type Hoteles []Hotel
