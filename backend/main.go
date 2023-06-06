package main

import (
	app "HotelArquiSoft/App"
	"HotelArquiSoft/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
