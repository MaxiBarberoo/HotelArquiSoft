package model

type Amenitie struct {
	ID     int    `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(100);not null; unique"`
}

type Amenities []Amenitie
