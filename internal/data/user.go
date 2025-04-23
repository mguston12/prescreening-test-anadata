package data

import (
	"Screening-Test-Anagata/go-project/internal/entity"

	"gorm.io/gorm"
)

type UserData interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type userData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) UserData {
	return &userData{db}
}

func (r *userData) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userData) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userData) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userData) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userData) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userData) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
