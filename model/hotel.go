package model

type Hotel struct {
	ID 			int 	`gorm:"primaryKey"`
	Nombre		string	`gorm:"type:varchar(100);not null"`
	CantHab 	int 	`gorm:"type:int;not null"`
	CantHabDis	int 	`gorm:"type:int"`
}

type Hotels []Hotel