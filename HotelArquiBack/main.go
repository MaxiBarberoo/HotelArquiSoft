package main

import (
	app "HotelArquiSoft/HotelArquiBack/App"
	"HotelArquiSoft/HotelArquiBack/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
