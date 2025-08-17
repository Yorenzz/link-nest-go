package service

import (
	"errors"
	"fmt"
	"log"

	"link-nest/internal/models"
	"link-nest/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	RegisterUser(username, email, password string) (*models.User, error)
	LoginUser(email, password string) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// RegisterUser handles user registration
func (s *userService) RegisterUser(username, email, password string) (*models.User, error) {
	// Check if user already exists by email
	_, err := s.userRepo.GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking existing user by email: %w", err)
	}

	// Check if username already exists
	_, err = s.userRepo.GetUserByUsername(username)
	if err == nil {
		return nil, errors.New("user with this username already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking existing user by username: %w", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	// Create user in database
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	log.Printf("User registered: %s", user.Email)
	return user, nil
}

// LoginUser handles user login
func (s *userService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	log.Printf("User logged in: %s", user.Email)
	return user, nil
}
