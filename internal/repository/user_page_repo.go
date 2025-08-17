package repository

import (
	"link-nest/internal/models"

	"gorm.io/gorm"
)

// UserPageRepository defines the interface for user page data operations
type UserPageRepository interface {
	CreateUserPage(userPage *models.UserPage) error
	GetUserPageByID(pageID uint) (*models.UserPage, error)
	GetUserPageBySlug(slug string) (*models.UserPage, error)
	GetActiveUserPageByUserID(userID uint) (*models.UserPage, error)
	GetUserPagesByUserID(userID uint) ([]models.UserPage, error)
}

type userPageRepository struct {
	db *gorm.DB
}

// NewUserPageRepository creates a new UserPageRepository
func NewUserPageRepository(db *gorm.DB) UserPageRepository {
	return &userPageRepository{db: db}
}

// CreateUserPage creates a new user page in the database
func (r *userPageRepository) CreateUserPage(userPage *models.UserPage) error {
	return r.db.Create(userPage).Error
}

// GetUserPageByID retrieves a user page by its ID
func (r *userPageRepository) GetUserPageByID(pageID uint) (*models.UserPage, error) {
	var userPage models.UserPage
	if err := r.db.First(&userPage, pageID).Error; err != nil {
		return nil, err
	}
	return &userPage, nil
}

// GetUserPageBySlug retrieves a user page by its slug
func (r *userPageRepository) GetUserPageBySlug(slug string) (*models.UserPage, error) {
	var userPage models.UserPage
	if err := r.db.Where("page_slug = ?", slug).First(&userPage).Error; err != nil {
		return nil, err
	}
	return &userPage, nil
}

// GetActiveUserPageByUserID retrieves the active user page for a user
func (r *userPageRepository) GetActiveUserPageByUserID(userID uint) (*models.UserPage, error) {
	var userPage models.UserPage
	if err := r.db.Where("user_id = ? AND is_active = ?", userID, true).First(&userPage).Error; err != nil {
		return nil, err
	}
	return &userPage, nil
}

// GetUserPagesByUserID retrieves all user pages for a user
func (r *userPageRepository) GetUserPagesByUserID(userID uint) ([]models.UserPage, error) {
	var userPages []models.UserPage
	if err := r.db.Where("user_id = ?", userID).Find(&userPages).Error; err != nil {
		return nil, err
	}
	return userPages, nil
}
