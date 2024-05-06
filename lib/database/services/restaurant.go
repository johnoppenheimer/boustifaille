package services

import (
	"gorm.io/gorm"

	"github.com/johnoppenheimer/boustifaille/database/models"
)

type RestaurantService struct {
	db *gorm.DB
}

func NewRestaurantService(db *gorm.DB) *RestaurantService {
	return &RestaurantService{db}
}

func (s *RestaurantService) Create(restaurant *models.Restaurant) error {
	return s.db.Create(&restaurant).Error
}

func (s *RestaurantService) All() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	err := s.db.Preload("User").Find(&restaurants).Error

	return restaurants, err
}

func (s *RestaurantService) FindBySlug(slug string) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := s.db.Preload("User").Where("slug = ?", slug).First(&restaurant).Error
	return &restaurant, err
}

func (s *RestaurantService) FindInBounds(lat1, lon1, lat2, lon2 float64) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	err := s.db.Where("latitude BETWEEN ? AND ?", lat1, lat2).Where("longitude BETWEEN ? AND ?", lon1, lon2).Find(&restaurants).Error
	return restaurants, err
}

func (s *RestaurantService) Update(restaurant *models.Restaurant) error {
	return s.db.Save(&restaurant).Error
}

func (s *RestaurantService) Delete(restaurant *models.Restaurant) error {
	return s.db.Delete(&restaurant).Error
}
