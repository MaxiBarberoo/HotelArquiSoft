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
		"password":   userDto.Password,
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
