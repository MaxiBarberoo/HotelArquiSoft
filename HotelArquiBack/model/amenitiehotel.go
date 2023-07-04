package model

type AmenitieHotel struct {
	ID         int `gorm:"primaryKey"`
	AmenitieId int `gorm:"foreignKey"`
	HotelId    int `gorm:"foreignKey"`
}

type AmenitiesHotels []AmenitieHotel
