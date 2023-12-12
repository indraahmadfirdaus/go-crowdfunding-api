package user

import (
	"crowdfunding-api/src/kernel"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{db: kernel.DB}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	user := User{}
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) FindByID(ID int) (User, error) {
	user := User{}
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
