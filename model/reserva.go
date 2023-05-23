package model

type Reserva struct {
	ID 			int
	FechaIn		time.Time 	`gorm:type:"DATETIME"`
	FechaOut 	time.Time 	`gorm:type:"DATETIME"`
	UserId		int 		`gorm:"foreignKey"`
	HotelId		int 		`gorm:"foreignKey"`
}

type Reservas []Reserva