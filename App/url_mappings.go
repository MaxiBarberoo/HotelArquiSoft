package app

import (
	//userController "HotelArquiSoft/controller/user"
	reservaController "HotelArquiSoft/controller/reserva"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/reservas", reservaController.GetReservas)
	router.GET("")
	log.Info("Finishing mappings configurations")
}
