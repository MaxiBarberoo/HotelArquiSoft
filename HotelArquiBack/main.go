package main

import (
	app "HotelArquiSoft/HotelArquiBack/app"
	"HotelArquiSoft/HotelArquiBack/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
