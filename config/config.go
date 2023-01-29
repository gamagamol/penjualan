package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY=[]byte("asdasd21whhjkhkjasd")
type JWTCLAIM struct{
	Ussername 				string
	jwt.RegisteredClaims
}