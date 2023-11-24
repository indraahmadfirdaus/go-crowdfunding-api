package user

import (
	"crowdfunding-api/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	authService := auth.NewAuthService()
	token, err := authService.GenerateToken(u.ID)
	if err != nil {
		return
	}
	u.Token = token
	return
}
