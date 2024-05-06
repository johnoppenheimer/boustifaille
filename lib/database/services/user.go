package services

import (
	"github.com/johnoppenheimer/boustifaille/database/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) Create(name string) error {
	return s.db.Create(&models.User{Name: name}).Error
}

func (s *UserService) FindByID(id uint) (models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	return user, err
}

func (s *UserService) Update(user *models.User) error {
	return s.db.Save(&user).Error
}
