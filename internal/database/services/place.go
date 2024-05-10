package services

import (
	"gorm.io/gorm"

	"github.com/johnoppenheimer/boustifaille/internal/database/models"
)

type PlaceService struct {
	db *gorm.DB
}

func NewPlaceService(db *gorm.DB) *PlaceService {
	return &PlaceService{db}
}

func (s *PlaceService) Create(place *models.Place) error {
	return s.db.Create(&place).Error
}

func (s *PlaceService) All() ([]models.Place, error) {
	var places []models.Place

	err := s.db.Preload("User").Find(&places).Error

	return places, err
}

func (s *PlaceService) FindBySlug(slug string) (*models.Place, error) {
	var place models.Place
	err := s.db.Preload("User").Where("slug = ?", slug).First(&place).Error
	return &place, err
}

func (s *PlaceService) FindInBounds(lat1, lon1, lat2, lon2 float64) ([]models.Place, error) {
	var places []models.Place
	err := s.db.Where("latitude BETWEEN ? AND ?", lat1, lat2).Where("longitude BETWEEN ? AND ?", lon1, lon2).Find(&places).Error
	return places, err
}

func (s *PlaceService) Update(place *models.Place) error {
	return s.db.Save(&place).Error
}

func (s *PlaceService) Delete(place *models.Place) error {
	return s.db.Delete(&place).Error
}
