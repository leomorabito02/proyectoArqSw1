package dto

type HotelDto struct {
	Id                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Descripcion       string  `json:"descripcion"`
	Cant_habitaciones int     `json:"cant_habitaciones"`
	Valoracion        int     `json:"valoracion"`
	Precio            float32 `json:"precio"`
}

type HotelesDto []HotelDto
