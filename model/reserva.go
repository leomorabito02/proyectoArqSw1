package model

type Reserva struct {
	Id           int     `gorm:"primaryKey"`
	Fecha_desde  string  `gorm:"type:date(300);not null"`
	Fecha_hasta  string  `gorm:"type:date(300);not null"`
	Precio_total float32 `gorm:"type:decimal(400);unsigned;not null"`

	User   User `gorm:"foreignkey:UsuarioId"`
	UserId int

	Hotel   Hotel `gorm:"foreignKey:HotelId"`
	HotelId int
}

type Reservas []Reserva
