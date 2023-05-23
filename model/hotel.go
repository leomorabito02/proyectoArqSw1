package model

type Hotel struct {
	Id                int     `gorm:"primaryKey"`
	Nombre            string  `gorm:"type:varchar(30);not null"`
	Descripcion       string  `gorm:"type:text();not null"`
	Cant_habitaciones int     `gorm:"type:int(3);not null"`
	Valoracion        int     `gorm:"type:int(1);unsigned;not null"`
	Gym               bool    `gorm:"type:boolean();not null"`
	Wifi              bool    `gorm:"type:boolean();not null"`
	Estacionamiento   bool    `gorm:"type:boolean();not null"`
	Bide              bool    `gorm:"type:boolean();not null"`
	Pileta            bool    `gorm:"type:boolean();not null"`
	Precio            float32 `gorm:"type:decimal(8,2);unsigned;not null"`
}

type Hoteles []Hotel
