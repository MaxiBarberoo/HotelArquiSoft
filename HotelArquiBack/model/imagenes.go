package model

type Imagen struct {
	ID        int `gorm:"primary_key"`
	HotelId   int `gorm:"foreignKey"`
	Nombre    string
	Contenido []byte
}
type Imagenes []Imagen
