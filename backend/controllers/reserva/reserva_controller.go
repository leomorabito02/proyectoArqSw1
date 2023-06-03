package reservaController

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"repo/dto"
	service "repo/services"
)

func ReservaInsert(c *gin.Context) {
	var reservaDto dto.ReservaDto
	err := c.BindJSON(&reservaDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservaDto, er := service.ReservaService.InsertReserva(reservaDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

//Buscar orden por token

func GetReservasByIdUser(c *gin.Context) {

	var reservasDto dto.ReservasDto
	token := c.Param("idUser")
	reservasDto, err := service.ReservaService.GetReservasByIdUser(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservasDto)
}
