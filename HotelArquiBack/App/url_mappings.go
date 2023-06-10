package app

import (
	hotelController "HotelArquiSoft/HotelArquiBack/controller/hotel"
	//userController "HotelArquiSoft/HotelArquiBack/controller/user"
	reservaController "HotelArquiSoft/HotelArquiBack/controller/reserva"
	userController "HotelArquiSoft/HotelArquiBack/controller/user"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/reservas", reservaController.GetReservas)
	router.GET("/reserva/:rooms", reservaController.GetRooms)
	router.GET("/reserva/:id", reservaController.GetReservaById)
	router.GET("/reserva/:reservauser", reservaController.GetReservasByUser)
	router.GET("/users", userController.GetUsers)
	router.GET("/user/:email", userController.GetUserByEmail)
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user/:auth", userController.UserAuth)
	router.GET("/hotels", hotelController.GetHotels)
	router.GET("/hotel/:id", hotelController.GetHotelById)

	router.POST("/user", userController.UserInsert)
	router.POST("/reserva", reservaController.ReservaInsert)
	router.POST("/hotel", hotelController.HotelInsert)

	log.Info("Finishing mappings configurations")
}
