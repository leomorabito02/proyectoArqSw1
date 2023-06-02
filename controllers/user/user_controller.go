package userController

import (
	"net/http"
	"repo/dto"
	service "repo/services"
	"strconv"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("ID de usuario a cargar: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

//login

func UserLogin(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := service.UserService.LoginUser(loginDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, tokenDto)

}

/*
func InsertUser(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}
*/
