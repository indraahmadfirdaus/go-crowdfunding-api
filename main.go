package main

import (
	"crowdfunding-api/auth"
	"crowdfunding-api/handler"
	"crowdfunding-api/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
	api.POST("/upload-avatar", userHandler.UpdateAvatar)

	router.Run()

}
