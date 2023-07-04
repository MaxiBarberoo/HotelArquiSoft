package model

type User struct {
	ID        int    `gorm:"primaryKey"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(50);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Tipo      int    `gorm:"type:int"`
}

type Users []User
