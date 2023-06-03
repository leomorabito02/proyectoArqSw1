package reserva

//ORM traductor
import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"repo/model"
)

var Db *gorm.DB

func InsertReserva(reserva model.Reserva) model.Reserva {
	result := Db.Create(&reserva)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Reserva creada: ", reserva.Id)
	return reserva
}

// busuqueda por idUser
func GetReservasByIdUser(idUser int) model.Reservas {
	var reservas model.Reservas

	log.Debug("idUser: ", idUser)
	Db.Where("id_user = ?", idUser).Find(&reservas)
	log.Debug("Reserva: ", reservas)

	return reservas
}
