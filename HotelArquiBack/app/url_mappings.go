package app

import (
	amenitieController "HotelArquiSoft/HotelArquiBack/controller/amenitie"
	amenitieHotelController "HotelArquiSoft/HotelArquiBack/controller/amenitiehotel"
	hotelController "HotelArquiSoft/HotelArquiBack/controller/hotel"
	imagenesController "HotelArquiSoft/HotelArquiBack/controller/imagenes"
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
	router.GET("amenities", amenitieController.GetAmenities)
	router.GET("/amenities/:id", amenitieController.GetAmenitieById)
	router.GET("/amenitiehotel/:hotel_id", amenitieHotelController.SearchAmenitiesByHotel)
	router.GET("/imagenes/:hotel_id", imagenesController.GetImagenesByHotel)
	router.GET("/reservas/hotel/:hotel_id", reservaController.GetReservasByHotel)

	router.POST("/users/auth", userController.UserAuth)
	router.POST("/imagenes", imagenesController.InsertImagen)
	router.POST("/reservas/hotelsbyfecha", reservaController.GetHotelsByFecha)
	router.POST("/users", userController.UserInsert)
	router.POST("/reservas", reservaController.ReservaInsert)
	router.POST("/hotels", hotelController.HotelInsert)
	router.POST("/reservas/rooms", reservaController.GetRooms)
	router.POST("reservas/byfecha", reservaController.GetReservasByFecha)
	router.POST("amenitiehotel/assign", amenitieHotelController.AssignAmenitieToHotel)
	router.POST("/reservas/:id", reservaController.GetReservaById)
	router.POST("/reservas/hoteluser", reservaController.GetReservasByHotelAndUser)
	router.POST("/reservas/fechauser", reservaController.GetReservasByFechaAndUser)
	router.POST("/reservas/fechahotel", reservaController.GetReservasByHotelAndFecha)
	router.POST("/reservas/hotelfechauser", reservaController.GetReservasByHotelFechaAndUser)
	router.GET("/reservas/reservauser/:user_id", reservaController.GetReservasByUser)
	log.Info("Finishing mappings configurations")
}
