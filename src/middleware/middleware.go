package middleware

import (
	"crowdfunding-api/src/domain/auth"
	"crowdfunding-api/src/domain/user"
)

type middleware struct {
	userService user.Service
	authService auth.Service
}

func NewMiddleware() *middleware {
	userService := user.NewService()
	authService := auth.NewAuthService()

	return &middleware{
		userService: userService,
		authService: authService,
	}
}
