package dto

import "time"

//type ReservaDtoRequest struct {
//	Fecha_desde time.Time `json:"fecha_desde"`
//	Fecha_hasta time.Time `json:"fecha_hasta"`
//	IdHotel     int       `json:"id_hotel"`
//	IdUser      int       `json:"id_user"`
//}
//type ReservaDtoResponse struct {
//	Id           int       `json:"id"`
//	Fecha_desde  time.Time `json:"fecha_desde"`
//	Fecha_hasta  time.Time `json:"fecha_hasta"`
//	Precio_total float32   `json:"precio_total"`
//	IdHotel      int       `json:"id_hotel"`
//	IdUser       int       `json:"id_user"`
//}

type ReservaDto struct {
	Id           int       `json:"id"`
	Fecha_desde  time.Time `json:"fecha_desde"`
	Fecha_hasta  time.Time `json:"fecha_hasta"`
	Precio_total float32   `json:"precio_total"`
	IdHotel      int       `json:"id_hotel"`
	IdUser       int       `json:"id_user"`
}

type ReservasDto []ReservaDto
