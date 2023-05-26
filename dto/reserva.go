package dto

type ReservaDto struct {
	Id           int     `json:"id"`
	Fecha_desde  string  `json:"fecha_desde"`
	Fecha_hasta  string  `json:"fecha_hasta"`
	Precio_total float32 `json:"precio_total"`
	IdHotel      int     `json:"id_hotel"`
	IdUser       int     `json:"id_user"`
}

type ReservasDto []ReservaDto
