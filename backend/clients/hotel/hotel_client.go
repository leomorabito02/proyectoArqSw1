package hotel

//ORM traductor
import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
type reservas struct {
	Cantidad int `gorm:"type:int64"`
}

func GetHabitacionesDisponibles(hotelID int, totalHabitaciones int, fecha_desde time.Time, fecha_hasta time.Time) int {

	var cantidad int64
	log.Debug("FECHA DESDE: ", fecha_desde)
	log.Debug("FECHA HASTA: ", fecha_hasta)
	from := fecha_desde.Format(time.DateOnly)
	to := fecha_hasta.Format(time.DateOnly)
	txn := Db.Model(&model.Reservas{}).
		Where("id_hotel= ?", hotelID).
		Where(Db.
			Where(Db.Where(Db.Where("fecha_desde <= ?", from).Where("? <= fecha_hasta", from))).
			Or(Db.Where(Db.Where("fecha_desde <= ?", to).Where("? <= fecha_hasta", to))).
			Or(Db.Where(Db.Where("? <= fecha_desde", from).Where("fecha_hasta <= ?", to)))).
		Count(&cantidad)
	if txn.Error != nil {
		log.Error(txn.Error)
		return 0
	}
	log.Debug("CANTIDAD ", cantidad)
	//cantidadDeReservasEnLaFechaSeleecionada = "SELECT COUNT(*) from hotels h JOIN reservas r ON h.id = r.hotel_id WHERE h.hotel_id = hotelID AND r.fecha_desde <= fecha_desde AND r.fecha_hasta >= fechaHasta"
	return totalHabitaciones - int(cantidad)
}
