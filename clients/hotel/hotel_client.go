package hotel

//ORM traductor
import (
	"proyectoArqSw1/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
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

func GetHotelesByDisponibilidad(fecha_desde string, fecha_hasta string) model.Hoteles {

	var hoteles model.Hoteles

	var count int
	Db.Model(&model.reserva{}).
		Where("fecha_desde <= ?", fecha_desde).
		And("fecha_hasta >= ?", fecha_hasta).
		Count(&count).
		Group("hotel_id")
	//Find(&hoteles)

	log.Debug("Hoteles disponibles: ", hoteles)

	return hoteles
}
