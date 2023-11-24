package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token *jwt.Token) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET = []byte("5SRRCRET_KKKEY#!!")

func NewAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET)

	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
