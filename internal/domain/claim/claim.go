package domain

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
