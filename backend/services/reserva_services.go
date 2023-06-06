package services

import (
	"github.com/dgrijalva/jwt-go"
	hotelCliente "repo/clients/hotel"
	reservaCliente "repo/clients/reserva"
	"repo/dto"
	model2 "repo/model"
	e "repo/utils/errors"
	"time"
)

type reservaService struct{}

type reservaServiceInterface interface {
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasByIdUser(token string) (dto.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var reserva model2.Reserva

	reserva.Id = reservaDto.Id
	reserva.Fecha_desde = reservaDto.Fecha_desde
	reserva.Fecha_hasta = reservaDto.Fecha_hasta
	reserva.IdUser = reservaDto.IdUser
	reserva.IdHotel = reservaDto.IdHotel

	var hotel model2.Hotel
	var diferencia time.Duration
	var cantDias int32

	hotel = hotelCliente.GetHotelById(reserva.IdHotel)

	var habitacionesDisponibles int
	habitacionesDisponibles = hotelCliente.GetHabitacionesDisponibles(reserva.IdHotel, hotel.Cant_habitaciones, reserva.Fecha_desde, reserva.Fecha_hasta)

	if habitacionesDisponibles <= 0 {
		var reservaResponseDto dto.ReservaDto

		return reservaResponseDto, e.NewBadRequestApiError("No hay habitaciones disponibles")
	}
	diferencia = reserva.Fecha_hasta.Sub(reserva.Fecha_desde)
	cantDias = int32(diferencia.Hours() / 24)
	reserva.Precio_total = float32(int32(hotel.Precio) * cantDias) //a super chequear
	reserva = reservaCliente.InsertReserva(reserva)

	var reservaResponseDto dto.ReservaDto

	reservaResponseDto.Id = reserva.Id
	reservaResponseDto.Fecha_desde = reserva.Fecha_desde
	reservaResponseDto.Fecha_hasta = reserva.Fecha_hasta
	reservaResponseDto.Precio_total = reserva.Precio_total
	reservaResponseDto.IdUser = reserva.IdUser
	reservaResponseDto.IdHotel = reserva.IdHotel

	return reservaResponseDto, nil

}

func (s *reservaService) GetReservasByIdUser(token string) (dto.ReservasDto, e.ApiError) {
	var idUser float64
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, e.NewUnauthorizedApiError("Unauthorized")
		}
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}

	if !tkn.Valid {
		return nil, e.NewUnauthorizedApiError("Unauthorized")

	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {

		idUser = (claims["id_user"].(float64))

	} else {
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}
	var IdUserX int = int(idUser)
	var reservas model2.Reservas = reservaCliente.GetReservasByIdUser(IdUserX)
	var reservasDto dto.ReservasDto

	if len(reservas) == 0 {
		return reservasDto, e.NewBadRequestApiError("No se encontro reserva")
	}

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto //a chequear
		reservaDto.Id = reserva.Id
		reservaDto.Fecha_desde = reserva.Fecha_desde
		reservaDto.Fecha_hasta = reserva.Fecha_hasta
		reservaDto.Precio_total = reserva.Precio_total
		reservaDto.IdHotel = reserva.IdHotel
		reservaDto.IdUser = reserva.IdUser

		reservasDto = append(reservasDto, reservaDto)

	}

	return reservasDto, nil
}
