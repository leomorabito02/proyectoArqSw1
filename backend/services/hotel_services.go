package services

import (
	hotelCliente "repo/clients/hotel"
	dto2 "repo/dto"
	"repo/model"
	e "repo/utils/errors"
	"time"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto2.HotelDto, e.ApiError)
	GetHoteles() (dto2.HotelesDto, e.ApiError)
	GetHotelDisponibilidad(hotelID int, fecha_desde time.Time, fecha_hasta time.Time) (dto2.Disponibilidad, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto2.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelCliente.GetHotelById(id)
	var hotelDto dto2.HotelDto

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

func (s *hotelService) GetHoteles() (dto2.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelCliente.GetHoteles()
	var hotelesDto dto2.HotelesDto
	if len(hoteles) == 0 {
		return hotelesDto, e.NewBadRequestApiError("hotels not found")
	}
	for _, hotel := range hoteles {
		var hotelDto dto2.HotelDto

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

func (s *hotelService) GetHotelDisponibilidad(hotelID int, fecha_desde time.Time, fecha_hasta time.Time) (dto2.Disponibilidad, e.ApiError) {

	hotel := hotelCliente.GetHotelById(hotelID)
	cantHDisponibles := hotelCliente.GetHabitacionesDisponibles(hotelID, hotel.Cant_habitaciones, fecha_desde, fecha_hasta)

	return dto2.Disponibilidad{
		HabitacionesDisponibles: cantHDisponibles,
		DetalleHotel: dto2.HotelDto{
			Id:                hotel.Id,
			Nombre:            hotel.Nombre,
			Descripcion:       hotel.Descripcion,
			Cant_habitaciones: hotel.Cant_habitaciones,
			Valoracion:        hotel.Valoracion,
			Precio:            hotel.Precio,
		},
	}, nil

}
