package dto

type UserDto struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email string `json:"email"`
	Password string `json:"password"`
	Tipo bool `json:"tipo"`
	Dni int `json:"dni"`	
}

type UsersDto []UserDto