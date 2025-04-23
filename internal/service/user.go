package service

import (
	data "Screening-Test-Anagata/go-project/internal/data"
	"Screening-Test-Anagata/go-project/internal/entity"
)

type UserService interface {
	Register(user *entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	UpdateByID(id uint, update *entity.User) (*entity.User, error)
	Delete(id uint) error
}

type userService struct {
	data data.UserData
}

func NewUserService(data data.UserData) UserService {
	return &userService{data}
}

func (s *userService) Register(user *entity.User) error {
	return s.data.Create(user)
}

func (s *userService) GetAll() ([]entity.User, error) {
	return s.data.FindAll()
}

func (s *userService) GetByID(id uint) (*entity.User, error) {
	return s.data.FindByID(id)
}

func (s *userService) Update(user *entity.User) error {
	return s.data.Update(user)
}

func (s *userService) UpdateByID(id uint, update *entity.User) (*entity.User, error) {
	existing, err := s.data.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Name = update.Name
	existing.Email = update.Email

	err = s.data.Update(existing)
	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *userService) Delete(id uint) error {
	return s.data.Delete(id)
}

func (s *userService) GetByEmail(email string) (*entity.User, error) {
	return s.data.FindByEmail(email)
}
