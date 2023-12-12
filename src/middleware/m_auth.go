package middleware

import (
	"crowdfunding-api/src/domain/constant"
	"crowdfunding-api/src/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (m *middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			helper.AbortResponse(c, constant.Unauthorized, nil)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, "Bearer ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := m.authService.ValidateToken(tokenString)

		if err != nil {
			helper.AbortResponse(c, constant.Unauthorized, nil)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			helper.AbortResponse(c, constant.Unauthorized, nil)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := m.userService.GetUserById(userID)

		if err != nil {
			helper.AbortResponse(c, constant.Unauthorized, nil)
			return
		}

		c.Set("currentUser", user)
	}
}
