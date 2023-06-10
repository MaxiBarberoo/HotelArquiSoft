package services

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	userClient "HotelArquiSoft/HotelArquiBack/clients/user"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	GetUserByEmail(email string) (dto.UserDto, e.ApiError)
	UserAuth(userDto dto.UserDto) bool
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserById(id)
	var userDto dto.UserDto

	if user.ID == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.UserEmail = user.Email
	userDto.Tipo = user.Tipo
	userDto.Id = user.ID

	return userDto, nil
}

func (s *userService) GetUserByEmail(email string) (dto.UserDto, e.ApiError) {
	var user model.User = userClient.GetUserByEmail(email)
	var userDto dto.UserDto

	if user.Email == "" {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.Tipo = user.Tipo
	userDto.Id = user.ID
	userDto.UserEmail = user.Email
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userClient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.Id = user.ID
		userDto.Tipo = user.Tipo

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Password = userDto.Password
	user.Email = userDto.UserEmail
	user.Tipo = userDto.Tipo

	user = userClient.InsertUser(user)

	userDto.Id = user.ID

	return userDto, nil
}

func (s *userService) UserAuth(userDto dto.UserDto) bool {

	user := userClient.GetUserByEmail(userDto.UserEmail)

	if user.Password != userDto.Password {
		return false
	}

	return true
}
