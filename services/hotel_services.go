package services

import (
	hotelCliente "repo/clients/hotel"
	"repo/dto"
	"repo/model"
	e "repo/utils/errors"
	"time"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	GetHotelesByDisponibilidad(fecha_desde time.Time, fecha_hasta time.Time) (dto.HotelesDto, e.ApiError)
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

		hotelesDto = append(hotelesDto, hotelDto) //append guarda la suma de "hotel"
	}

	return hotelesDto, nil
}

func (s *hotelService) GetHotelesByDisponibilidad(fecha_desde time.Time, fecha_hasta time.Time) (dto.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelCliente.GetHotelesByDisponibilidad(fecha_desde, fecha_hasta)
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

		hotelesDto = append(hotelesDto, hotelDto) //append guarda la suma de "hotel"
	}

	return hotelesDto, nil

}
