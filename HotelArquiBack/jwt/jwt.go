package jwtToken

import (
	"HotelArquiSoft/HotelArquiBack/dto"

	"github.com/dgrijalva/jwt-go"
)

func GenerateUserToken(userDto dto.UserDto) (string, error) {
	claims := jwt.MapClaims{
		"id":         userDto.Id,
		"name":       userDto.FirstName,
		"last_name":  userDto.LastName,
		"user_email": userDto.UserEmail,
		"tipo":       userDto.Tipo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secret := "secreto"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateHotelToken(hotelDto dto.HotelDto) (string, error) {
	claims := jwt.MapClaims{
		"id":                hotelDto.Id,
		"name":              hotelDto.Name,
		"cant_habitaciones": hotelDto.CantHabitaciones,
		"desc":              hotelDto.Desc,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateReservaToken(reservaDto dto.ReservaDto) (string, error) {
	claims := jwt.MapClaims{
		"id":            reservaDto.Id,
		"user_id":       reservaDto.UserId,
		"hotel_id":      reservaDto.HotelId,
		"fecha_ingreso": reservaDto.FechaIngreso,
		"fecha_egreso":  reservaDto.FechaEgreso,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateAmenitieToken(amenitieDto dto.AmenitieDto) (string, error) {
	claims := jwt.MapClaims{
		"id":   amenitieDto.Id,
		"tipo": amenitieDto.Tipo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateAmentieHotelToken(amenitieHotelDto dto.AmenitieHotelDto) (string, error) {
	claims := jwt.MapClaims{
		"id":          amenitieHotelDto.Id,
		"amenitie_id": amenitieHotelDto.AmenitieId,
		"hotel_id":    amenitieHotelDto.HotelId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
func GenerateImageToken(ImageDto dto.ImageDTO) (string, error) {
	claims := jwt.MapClaims{
		"id":        ImageDto.ID,
		"nombre":    ImageDto.Nombre,
		"contenido": ImageDto.Contenido,
		"hotel_id":  ImageDto.HotelId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
