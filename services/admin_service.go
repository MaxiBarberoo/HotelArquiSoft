package services

import (
	e "HotelArquiSoft/Utils"
	adminClient "HotelArquiSoft/clients/admin"
	"HotelArquiSoft/dto"
	"HotelArquiSoft/model"
)

type adminService struct{}

type adminServiceInterface interface {
	GetAdminById(id int) (dto.AdminDto, e.ApiError)
	GetAdmins() (dto.AdminsDto, e.ApiError)
	InsertAdmin(adminDto dto.AdminDto) (dto.AdminDto, e.ApiError)
}

var (
	AdminService adminServiceInterface
)

func init() {
	AdminService = &adminService{}
}

func (s *adminService) GetAdminById(id int) (dto.AdminDto, e.ApiError) {

	var admin model.Admin = adminClient.GetAdminById(id)
	var adminDto dto.AdminDto

	if admin.ID == 0 {
		return adminDto, e.NewBadRequestApiError("user not found")
	}

	adminDto.FirstName = admin.FirstName
	adminDto.LastName = admin.LastName

	return adminDto, nil
}

func (s *adminService) GetAdmins() (dto.AdminsDto, e.ApiError) {

	var admin model.Admins = adminClient.GetAdmins()
	var adminsDto dto.AdminsDto

	for _, admin := range admin {
		var adminDto dto.AdminDto
		adminDto.FirstName = admin.FirstName
		adminDto.LastName = admin.LastName
		adminDto.Id = admin.ID

		adminsDto = append(adminsDto, adminDto)
	}

	return adminsDto, nil
}

func (s *adminService) InsertAdmin(adminDto dto.AdminDto) (dto.AdminDto, e.ApiError) {

	var admin model.Admin

	admin.FirstName = adminDto.FirstName
	admin.LastName = adminDto.LastName
	admin.Password = adminDto.Password

	adminDto.Id = admin.ID

	admin = adminClient.InsertAdmin(admin)

	return adminDto, nil
}
