package repositories

import (
	"backend-journey/models"

	"gorm.io/gorm"
)

type BookmarkRepository interface {
	CreateBookmark(User models.Bookmark) (models.Bookmark, error)
	FindBookmarks() ([]models.Bookmark, error)
	GetBookmark(ID int) (models.Bookmark, error)
	UpdateBookmark(User models.Bookmark) (models.Bookmark, error)
	DeleteBookmark(User models.Bookmark) (models.Bookmark, error)
}

func RepositoryBookmark(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBookmark(bookmark models.Bookmark) (models.Bookmark, error) {
	err := r.db.Preload("User").Preload("Journey.User").Create(&bookmark).Error

	return bookmark, err
}

func (r *repository) FindBookmarks() ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	err := r.db.Preload("User").Preload("Journey.User").Find(&bookmarks).Error

	return bookmarks, err
}

func (r *repository) GetBookmark(ID int) (models.Bookmark, error) {
	var bookmark models.Bookmark
	err := r.db.Preload("User").Preload("Journey.User").First(&bookmark, ID).Error

	return bookmark, err
}

func (r *repository) UpdateBookmark(bookmark models.Bookmark) (models.Bookmark, error) {
	err := r.db.Preload("User").Preload("Journey.User").Save(&bookmark).Error

	return bookmark, err
}

func (r *repository) DeleteBookmark(bookmark models.Bookmark) (models.Bookmark, error) {
	err := r.db.Preload("User").Preload("Journey.User").Delete(&bookmark).Error

	return bookmark, err
}
