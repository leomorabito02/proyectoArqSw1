package services

import (
	"repo/dto"
	"repo/model"
	e "repo/utils/errors"

	userCliente "repo/clients/user"

	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"

	log "github.com/sirupsen/logrus"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Id = user.Id
	userDto.Nombre = user.Nombre
	userDto.Apellido = user.Apellido
	userDto.Email = user.Email
	userDto.Password = user.Password
	userDto.Tipo = user.Tipo
	userDto.Dni = user.Dni

	return userDto, nil
}

//login

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) {

	log.Debug(loginDto) //para registrar el contenido de loginDto
	var user model.User = userCliente.GetUserByEmail(loginDto.Email)

	var tokenDto dto.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	//pasamos password como slice de bytes
	//hashea con md5.sum
	var pswMd5 = md5.Sum([]byte(loginDto.Password))
	//convertir a cadena hexadecimal
	pswMd5String := hex.EncodeToString(pswMd5[:])

	if pswMd5String == user.Password {
		//se firma el token para verificar autenticidad
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id

		return tokenDto, nil
	} else {
		return tokenDto, e.NewBadRequestApiError("contraseña incorrecta")
	}

}
func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.Id = userDto.Id
	user.Nombre = userDto.Nombre
	user.Apellido = userDto.Apellido
	user.Email = userDto.Email
	user.Password = userDto.Password
	user.Tipo = userDto.Tipo
	user.Dni = userDto.Dni

	user = userCliente.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}
