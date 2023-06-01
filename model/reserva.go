package model

import "time"

type Reserva struct {
	Id           int       `gorm:"primaryKey"`
	Fecha_desde  time.Time `gorm:"type:date;not null"`
	Fecha_hasta  time.Time `gorm:"type:date;not null"`
	Precio_total float32   `gorm:"type:decimal;unsigned;not null"`
	IdHotel      int       `gorm:"type:integer;not null"`
	IdUser       int       `gorm:"type:integer;not null"`

	// ESTO ESTA BIEN? es necesatrio?
	//User   User `gorm:"foreignkey:UsuarioId"`
	//UserId int

	//Hotel   Hotel `gorm:"foreignKey:HotelId"`
	//HotelId int
}

type Reservas []Reserva
