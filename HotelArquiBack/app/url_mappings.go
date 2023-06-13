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
	router.GET("/users", userController.GetUsers)
	router.GET("/users/email", userController.GetUserByEmail)
	router.GET("/users/:id", userController.GetUserById)
	router.GET("/hotels", hotelController.GetHotels)
	router.GET("/hotels/:id", hotelController.GetHotelById)

	router.POST("/users/auth", userController.UserAuth)
	router.POST("/users", userController.UserInsert)
	router.POST("/reservas", reservaController.ReservaInsert)
	router.POST("/hotels", hotelController.HotelInsert)
	router.POST("/reservas/rooms", reservaController.GetRooms)
	router.POST("/reservas/:id", reservaController.GetReservaById)
	router.GET("/reservas/reservauser/:user_id", reservaController.GetReservasByUser)
	log.Info("Finishing mappings configurations")
}
