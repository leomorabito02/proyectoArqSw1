package services

import (
	hotelCliente "proyectoArqSw1/clients/hotel"
	"proyectoArqSw1/dto"
	"proyectoArqSw1/model"
	e "proyectoArqSw1/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	GetHotelesByDisponibilidad(fecha_desde string, fecha_hasta string) (dto.HotelesDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelCliente.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}
	hotelDto.Id = hotel.Id
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Cant_habitaciones = hotel.Cant_habitaciones
	hotelDto.Valoracion = hotel.Valoracion
	hotelDto.Precio = hotel.Precio

	return hotelDto, nil
}

func (s *hotelService) GetHoteles() (dto.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelCliente.GetHoteles()
	var hotelesDto dto.HotelesDto
	if len(hoteles) == 0 {
		return hotelesDto, e.NewBadRequestApiError("hotels not found")
	}
	for _, hotel := range hoteles {
		var hotelDto dto.HotelDto

		hotelDto.Id = hotel.Id
		hotelDto.Nombre = hotel.Nombre
		hotelDto.Descripcion = hotel.Descripcion
		hotelDto.Cant_habitaciones = hotel.Cant_habitaciones
		hotelDto.Valoracion = hotel.Valoracion
		hotelDto.Precio = hotel.Precio

		hotelesDto = append(hotelesDto, hotelesDto) //append guarda la suma de "hotel"
	}

	return hotelesDto, nil
}

func (s *hotelService) GetHotelesByDisponibilidad(fecha_desde string, fecha_hasta string) (dto.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelCliente.GetHotelesByDisponibilidad()
	var hotelesDto dto.HotelesDto
	if len(hoteles) == 0 {
		return hotelesDto, e.NewBadRequestApiError("No hay hoteles disponibles para la fecha")
	}
	for _, hotel := range hoteles {
		var hotelDto dto.HotelDto

		hotelDto.Id = hotel.Id
		hotelDto.Nombre = hotel.Nombre
		hotelDto.Descripcion = hotel.Descripcion
		hotelDto.Cant_habitaciones = hotel.Cant_habitaciones
		hotelDto.Valoracion = hotel.Valoracion
		hotelDto.Precio = hotel.Precio

		hotelesDto = append(hotelesDto, hotelesDto) //append guarda la suma de "hotel"
	}

	return hotelesDto, nil

}
