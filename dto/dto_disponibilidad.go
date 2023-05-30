package dto

type Disponibilidad struct {
	HabitacionesDisponibles int      `json:"habitaciones_disponibles"`
	DetalleHotel            HotelDto `json:"detalle_hotel"`
}
