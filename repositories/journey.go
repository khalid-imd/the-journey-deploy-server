package repositories

import (
	"backend-journey/models"

	"gorm.io/gorm"
)

type JourneyRepository interface {
	FindJourneys() ([]models.Journey, error)
	GetJourney(ID int) (models.Journey, error)
	CreateJourney(journey models.Journey) (models.Journey, error)
}

type repositoryJourney struct {
	db *gorm.DB
}

func RepositoryJourney(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateJourney(journey models.Journey) (models.Journey, error) {
	err := r.db.Preload("User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Create(&journey).Error

	return journey, err
}

func (r *repository) FindJourneys() ([]models.Journey, error) {
	var journeys []models.Journey
	err := r.db.Preload("User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Find(&journeys).Error

	return journeys, err
}

func (r *repository) GetJourney(ID int) (models.Journey, error) {
	var journey models.Journey
	err := r.db.Preload("User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").First(&journey, ID).Error

	return journey, err
}

func (r *repository) UpdateJourney(journey models.Journey, ID int) (models.Journey, error) {
	err := r.db.Save(&journey).Error

	return journey, err
}

func (r *repository) DeleteJourney(journey models.Journey, ID int) (models.Journey, error) {
	err := r.db.Delete(&journey).Error

	return journey, err
}
