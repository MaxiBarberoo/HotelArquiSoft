package model

type User struct {
	ID        int    `gorm:"primaryKey"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(20);not null"`
	Tipo      string `gorm:"type:varchar(10);not null"`
}

type Users []User
