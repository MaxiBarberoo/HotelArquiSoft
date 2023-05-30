package controller

import (
	"HotelArquiSoft/dto"
	service "HotelArquiSoft/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAdmin(c *gin.Context) {
	log.Debug("Admin id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdmins(c *gin.Context) {
	var adminsDto dto.AdminsDto
	adminsDto, err := service.AdminService.GetAdmins()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, adminsDto)
}

func AdminInsert(c *gin.Context) {
	var adminDto dto.AdminDto
	err := c.BindJSON(&adminDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	adminDto, er := service.AdminService.InsertAdmin(adminDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, adminDto)
}
