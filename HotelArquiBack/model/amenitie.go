package model

type Amenitie struct {
	ID   int    `gorm:"primaryKey"`
	Tipo string `gorm:"type:varchar(100);not null; unique"`
}

type Amenities []Amenitie
