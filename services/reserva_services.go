package services

import (
	reservaCliente "proyectoArqSw1/clients/reserva"
	"proyectoArqSw1/dto"
	"proyectoArqSw1/model"
	e "proyectoArqSw1/utils/errors"
)

type reservaService struct{}

type reservaServiceInterface interface {
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservaByIdUser(token string) (dto.ReservaDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva

	reservaDto.Id = reserva.Id
	reserva.Fecha_desde = reservaDto.Fecha_desde
	reserva.Fecha_hasta = reservaDto.Fecha_hasta
	reserva.IdUser = reservaDto.IdUser
	reserva.IdHotel = reservaDto.IdHotel

	var hotelDto dto.HotelDto
	var diferencia, cantDias int32

	hotelDto, _ = GetHotelById(reserva.IdHotel)
	diferencia = reserva.Fecha_hasta.Sub(reserva.Fecha_desde)
	cantDias = int32(int(diferencia.Hours() / 24))
	reserva.Precio_total = hotelDto.Precio * cantDias
	reserva = reservaCliente.InsertReserva(reserva)

	return reservaDto, nil
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
	var reservas model.Reservas = reservaCliente.GetReservasByIdUser(IdUserX)
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
