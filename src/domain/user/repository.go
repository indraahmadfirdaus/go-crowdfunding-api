package user

import (
	"crowdfunding-api/src/kernel"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Save(user User) (User, error) {
	err := kernel.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	user := User{}
	err := kernel.DB.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) FindByID(ID int) (User, error) {
	user := User{}
	err := kernel.DB.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) Update(user User) (User, error) {
	err := kernel.DB.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
