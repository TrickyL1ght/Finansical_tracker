package user

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	CrateUser(user User) error
	GetUser(login, password string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) CrateUser(user User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return fmt.Errorf("User Register Error: %v", err)
	}
	return nil
}

func (r *repository) GetUser(login, password string) (User, error) {
	var user User
	if err := r.db.Where("login=? and password=?", login, password).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("This user not found: %v", err)
	}
	return user, nil
}
