package jwt

import "time"

type Tokens struct {
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
	Exp          time.Duration `json:"-"`
}
