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

	reserva.Fecha_desde = reservaDto.Fecha_desde

	reserva.Fecha_hasta = reservaDto.Fecha_hasta

	//reserva.IdUser = reservaDto.IdUser

	//reserva.IdHotel = reservaDto.IdHotel

	reserva = reservaCliente.InsertReserva(reserva)

	reservaDto.Id = reserva.Id

	var precio_total float32

}
