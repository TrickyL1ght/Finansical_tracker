package repository

import (
	"awesomeProject/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type UserGORM struct {
	db *gorm.DB
}

func NewUserGORM(db *gorm.DB) *UserGORM {
	return &UserGORM{db: db}
}

func (r *UserGORM) CrateUser(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return fmt.Errorf("User Register Error: %v", err)
	}
	return nil
}

func (r *UserGORM) GetUser(login, password string) (model.User, error) {
	var user model.User
	if err := r.db.Where("login=? and password=?", login, password).First(&user).Error; err != nil {
		return model.User{}, fmt.Errorf("This user not found: %v", err)
	}
	return user, nil
}
