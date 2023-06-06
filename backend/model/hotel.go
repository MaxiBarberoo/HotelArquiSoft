package model

type Hotel struct {
	ID      int    `gorm:"primaryKey"`
	Nombre  string `gorm:"type:varchar(100);not null"`
	CantHab int    `gorm:"type:int;not null"`
}

type Hotels []Hotel
