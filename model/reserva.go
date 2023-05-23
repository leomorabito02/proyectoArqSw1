package model

type Reserva struct {
	Id           int     `gorm:"primaryKey"`
	Fecha_desde  string  `gorm:"type:date();not null"`
	Fecha_hasta  string  `gorm:"type:date();not null"`
	Precio_total float32 `gorm:"type:decimal(8,2);unsigned;not null"`

	Usuario   Usuario `gorm:"foreignkey:UsuarioId"`
	UsuarioId int

	Hotel   Hotel `gorm:"foreignKey:HotelId"`
	HotelId int
}

type Reservas []Reserva
