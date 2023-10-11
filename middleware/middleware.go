package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Id    string `json:"id"`
	Level int    `json:"level"`
	jwt.StandardClaims
}

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: []byte("secret key handokowae.my.id"),
})
