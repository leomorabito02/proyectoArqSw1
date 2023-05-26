package app

import (
	log "github.com/sirupsen/logrus"
)

// para el final dockerizar hacer git pull y docker compose para levantar todo de una
// y testing, md5 para la contraseña no tnenerlo en texto plano, hasheo
func mapUrls() {
	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.POST("/login", userController.UserLogin)

	//Hotel Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", productController.GetHoteles) //agregar

	//Order Mapping
	router.POST("/reserva", reservaController.ReservaInsert)

	router.GET("/reservaUser/:idUser", reservaController.GetReservasByIdUser)

	log.Info("Listo el mapeo de configuraciones :)")
}
