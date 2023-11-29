package main

import (
	"crowdfunding-api/auth"
	"crowdfunding-api/constant"
	"crowdfunding-api/handler"
	"crowdfunding-api/helper"
	"crowdfunding-api/user"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=crowdfunding-db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authService := auth.NewAuthService()

	router := gin.Default()
	api := router.Group("/api/v1")

	fmt.Println(authService.GenerateToken(100))

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailAvailability)
	api.POST("/upload-avatar", authMiddleware(authService, userService), userHandler.UpdateAvatar)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			helper.AbortResponse(c, constant.Unauthorized, nil)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, "Bearer")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

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

		user, err := userService.GetUserById(userID)

		if err != nil {
			helper.AbortResponse(c, constant.Unauthorized, nil)
			return
		}

		c.Set("currentUser", user)
	}
}
