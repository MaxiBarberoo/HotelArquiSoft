package model

type Imagen struct {
	ID        int    `gorm:"primary_key"`
	HotelId   int    `gorm:"foreignKey"`
	Nombre    string `gorm:"type:varchar(500);not null"`
	Contenido []byte `gorm:"type:longblob"`
}
type Imagenes []Imagen
