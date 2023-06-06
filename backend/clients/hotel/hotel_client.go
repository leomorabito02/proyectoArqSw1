package hotel

//ORM traductor
import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"repo/model"
	"time"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHoteles() model.Hoteles {
	var hoteles model.Hoteles
	Db.Find(&hoteles)

	log.Debug("Hoteles: ", hoteles)

	return hoteles
}

/*
func GetHabitacionesDisponibles(hotelID int, totalHabitaciones int, fecha_desde time.Time, fecha_hasta time.Time) int {

	var cantReservas int
	Db.Model(&model.Hotel{}).
		Joins("join reservas on hoteles.id = reservas.hotel_id").
		Where("hoteles.id= ?", hotelID).
		Where("reservas.fecha_desde >= ?", fecha_desde).
		Where("reservas.fecha_hasta <= ?", fecha_hasta).
		Count(&cantReservas)

	//cantidadDeReservasEnLaFechaSeleecionada = "SELECT COUNT(*) from hotels h JOIN reservas r ON h.id = r.hotel_id WHERE h.hotel_id = hotelID AND r.fecha_desde <= fecha_desde AND r.fecha_hasta >= fechaHasta"

	return totalHabitaciones - cantReservas
}
*/

func GetHabitacionesDisponibles(hotelID int, totalHabitaciones int, fecha_desde time.Time, fecha_hasta time.Time) int {

	var cantReservas int
	Db.Model(&model.Hotel{}).
		Joins("join reservas on hotels.id = reservas.hotel_id").
		Where("hotels.id= ? AND reservas.fecha_desde >= ? AND reservas.fecha_hasta <= ?", hotelID, fecha_desde, fecha_hasta).
		Count(&cantReservas)

	//cantidadDeReservasEnLaFechaSeleecionada = "SELECT COUNT(*) from hotels h JOIN reservas r ON h.id = r.hotel_id WHERE h.hotel_id = hotelID AND r.fecha_desde <= fecha_desde AND r.fecha_hasta >= fechaHasta"

	return totalHabitaciones - cantReservas
}
