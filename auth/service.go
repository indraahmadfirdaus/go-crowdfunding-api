package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userId int) (string, error)
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
