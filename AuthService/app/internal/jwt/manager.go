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
	GenerateToken(int, string, string) (string, error)
	GenerateRefreshToken() (string, error)
	ParseToken(string) (string, string, error)
}

type Manager struct {
	key []byte
}

func NewTokenManager(key string) (*Manager, error) {
	if key == "" {
		return nil, errors.New("empty token manager key")
	}

	return &Manager{
		key: []byte(key),
	}, nil
}

func (m *Manager) GenerateToken(id int, email string, role string) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  role,
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Subject:   strconv.Itoa(id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(m.key)
	if err != nil {
		return "", err
	}
	return ss, nil
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

func (m *Manager) ParseToken(token string) (string, string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(m.key), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("error get claims from token")
	}

	return claims["sub"].(string), claims["aud"].(string), nil
}
