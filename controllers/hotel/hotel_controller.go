package hotelController

import (
	"net/http"
	"repo/dto"
	service "repo/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHotelById(c *gin.Context) {
	log.Debug("id de Hotel a cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHoteles(c *gin.Context) {
	var hotelesDto dto.HotelesDto
	hotelesDto, err := service.HotelService.GetHoteles()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelesDto)
}

func GetHotelesByDisponibilidad(c *gin.Context) {
	fechaDesdeStr := c.Query("fecha_desde")
	fechaHastaStr := c.Query("fecha_hasta")
	// hace falta convertir tipo de dato de fecha¿?
	var hotelesDto dto.HotelesDto

	fechaDesde, err := time.Parse("2006-01-02", fechaDesdeStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato fecha_desde inválido"})
		return
	}

	fechaHasta, err := time.Parse("2006-01-02", fechaHastaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato fecha_hasta inválido"})
		return
	}

	hotelesDto, err = service.HotelService.GetHotelesByDisponibilidad(fechaDesde, fechaHasta)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelesDto)
}
