package hotel

//ORM traductor
import (
	"repo/model"

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
	/*
		Db.Select("hotel_id", "cant_habitaciones").Find(&hoteles)

		Db.Model(&model.reserva{}).
			Where("fecha_desde <= ?", fecha_desde).
			And("fecha_hasta >= ?", fecha_hasta).
			Count(&count).
			Group("hotel_id")
		//Find(&hoteles)

		Db.Model(&model.hotel{}).Select("users.name, emails.email").
			Joins("left join emails on emails.user_id = users.id").
			Scan(&result{})
		// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

		log.Debug("Hoteles disponibles: ", hoteles)
	*/
	Db.Joins("JOIN hoteles ON hoteles.id = reserva.hotel_id").
		Select("hotel.id, hotel.Cant_habitaciones, reserva.Fecha_desde, reserva.Fecha_hasta").
		Find(&reservas)

	return hoteles
}
