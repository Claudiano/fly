package models

import jwt "github.com/dgrijalva/jwt-go"

type Claim struct {
	Passageiro `json:"passageiro"`
	jwt.StandardClaims
}
