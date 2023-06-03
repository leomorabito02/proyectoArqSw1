package model

type Foto struct {
	Id      int    `gorm:"primaryKey"`
	Url     string `gorm:"type:varchar(400);not null"`
	Hotel   Hotel  `gorm:"foreignKey:HotelId"`
	HotelId int
}

type Fotos []Foto
