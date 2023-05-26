package dto

import "time"

type ReservaDto struct {
	Id           int       `json:"id"`
	Fecha_desde  time.Time `json:"fecha_desde"`
	Fecha_hasta  time.Time `json:"fecha_hasta"`
	Precio_total float32   `json:"precio_total"`
	IdHotel      int       `json:"id_hotel"`
	IdUser       int       `json:"id_user"`
}

type ReservasDto []ReservaDto
