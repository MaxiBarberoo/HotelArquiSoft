package dto

type UserDto struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	UserEmail string `json:"user_email"`
	Password  string `json:"password"`
}

type UsersDto []UserDto
