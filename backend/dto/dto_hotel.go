package dto

type HotelDto struct {
	Id                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Descripcion       string  `json:"descripcion"`
	Cant_habitaciones int     `json:"cant_habitaciones"`
	Valoracion        int     `json:"valoracion"`
	Precio            float32 `json:"precio"`
	Gym               bool    `json:"gym"`
	Wifi              bool    `json:"wifi"`
	Estacionamiento   bool    `gorm:"estacionamiento"`
	Bide              bool    `gorm:"bide"`
	Pileta            bool    `gorm:"pileta"`
}

type HotelesDto []HotelDto
