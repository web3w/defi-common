package token

import "github.com/dgrijalva/jwt-go"

type Claims interface {
	jwt.Claims
	Id() string
	Trader() string
	Chain() string
	UserID() int64
}

type Generator interface {
	MarshaledVerifyKey() string
	MarshaledSignKey() string
	GenerateToken(claims Claims) (string, error)
}

type Verifier interface {
	VerifyToken(string) (Claims, error)
}
