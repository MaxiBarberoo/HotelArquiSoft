package services

import (
	e "HotelArquiSoft/Utils"
	userClient "HotelArquiSoft/clients/user"
	"HotelArquiSoft/dto"
	"HotelArquiSoft/model"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
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

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Password = userDto.Password

	userDto.Id = user.ID

	user = userClient.InsertUser(user)

	return userDto, nil
}
