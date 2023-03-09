package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthPayload struct {
	UserId    uint      `json:"userId"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (p AuthPayload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.New("Invalid token")
	}
	return nil
}

func GenerateToken(userId uint) (string, error) {
	payload := AuthPayload{userId, time.Now(), time.Now().Add(time.Hour * 720)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(token string) (*AuthPayload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	}

	var payload AuthPayload

	_, err := jwt.ParseWithClaims(token, &payload, keyFunc)

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	return &payload, nil
}
