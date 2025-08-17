package service

import (
	"fmt"
	"log"

	"link-nest/internal/models"
	"link-nest/internal/repository"
)

// UserPageService defines the interface for user page-related business logic
type UserPageService interface {
	CreateUserPage(userID uint, slug, title string, templateID *uint) (*models.UserPage, error)
	GetUserPageBySlug(slug string) (*models.UserPage, error)
	GetActiveUserPageByUserID(userID uint) (*models.UserPage, error)
}

type userPageService struct {
	userPageRepo repository.UserPageRepository
}

// NewUserPageService creates a new UserPageService
func NewUserPageService(userPageRepo repository.UserPageRepository) UserPageService {
	return &userPageService{userPageRepo: userPageRepo}
}

// CreateUserPage handles user page creation
func (s *userPageService) CreateUserPage(userID uint, slug, title string, templateID *uint) (*models.UserPage, error) {
	userPage := &models.UserPage{
		UserID:     userID,
		PageSlug:   slug,
		PageTitle:  title,
		TemplateID: templateID,
		IsActive:   true,
	}

	// Create user page in database
	if err := s.userPageRepo.CreateUserPage(userPage); err != nil {
		return nil, fmt.Errorf("failed to create user page: %w", err)
	}

	log.Printf("User page created: %s for user %s", userPage.PageSlug, userPage.UserID)
	return userPage, nil
}

// GetUserPageBySlug retrieves a user page by slug
func (s *userPageService) GetUserPageBySlug(slug string) (*models.UserPage, error) {
	userPage, err := s.userPageRepo.GetUserPageBySlug(slug)
	if err != nil {
		return nil, fmt.Errorf("failed to get user page by slug: %w", err)
	}
	return userPage, nil
}

// GetActiveUserPageByUserID retrieves the active user page for a user
func (s *userPageService) GetActiveUserPageByUserID(userID uint) (*models.UserPage, error) {
	userPage, err := s.userPageRepo.GetActiveUserPageByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get active user page: %w", err)
	}
	return userPage, nil
}
