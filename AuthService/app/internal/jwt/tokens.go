package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	ProfileID int `json:"profileId"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
	Exp          time.Duration `json:"-"`
}
