package jwt

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	GenerateToken(id int, role string, profileId int) (string, error)
	GenerateRefreshToken() (string, error)
	ParseToken(token string) (*tokenClaims, error)
	GetExpTime() time.Duration
}

type Manager struct {
	key []byte
	exp time.Duration
}

func NewTokenManager(key string, exp time.Duration) (TokenManager, error) {
	if key == "" {
		return nil, errors.New("empty token manager key")
	}

	return &Manager{
		key: []byte(key),
		exp: exp,
	}, nil
}

func (m *Manager) GenerateToken(id int, role string, profileId int) (string, error) {
	claims := tokenClaims{
		profileId,
		jwt.StandardClaims{
			Audience:  role,
			ExpiresAt: time.Now().Add(time.Minute * m.exp).Unix(),
			Subject:   strconv.Itoa(id),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(m.key)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (m *Manager) GetExpTime() time.Duration {
	return time.Duration(time.Now().Add(m.exp).Unix())
}

func (m *Manager) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)

	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)

	_, err := random.Read(b)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *Manager) ParseToken(token string) (*tokenClaims, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(m.key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*tokenClaims)
	if !ok {
		return nil, fmt.Errorf("error get claims from token")
	}

	return claims, nil
}
