package service

import (
	"fmt"
	"log"

	"link-nest/internal/models"
	"link-nest/internal/repository"

	"gorm.io/datatypes"
)

// ModuleService defines the interface for module-related business logic
type ModuleService interface {
	CreateModule(pageID uint, moduleType string, title *string, order int, config datatypes.JSON) (*models.Module, error)
	GetModulesByPageID(pageID uint) ([]models.Module, error)
	GetEnabledModulesByPageID(pageID uint) ([]models.Module, error)
}

type moduleService struct {
	moduleRepo repository.ModuleRepository
}

// NewModuleService creates a new ModuleService
func NewModuleService(moduleRepo repository.ModuleRepository) ModuleService {
	return &moduleService{moduleRepo: moduleRepo}
}

// CreateModule handles module creation
func (s *moduleService) CreateModule(pageID uint, moduleType string, title *string, order int, config datatypes.JSON) (*models.Module, error) {
	module := &models.Module{
		PageID:      pageID,
		ModuleType:  moduleType,
		ModuleTitle: title,
		ModuleOrder: order,
		Config:      config,
		IsEnabled:   true,
	}

	// Create module in database
	if err := s.moduleRepo.CreateModule(module); err != nil {
		return nil, fmt.Errorf("failed to create module: %w", err)
	}

	log.Printf("Module created: %s for page %s", module.ModuleType, module.PageID)
	return module, nil
}

// GetModulesByPageID retrieves all modules for a page
func (s *moduleService) GetModulesByPageID(pageID uint) ([]models.Module, error) {
	modules, err := s.moduleRepo.GetModulesByPageID(pageID)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules by page ID: %w", err)
	}
	return modules, nil
}

// GetEnabledModulesByPageID retrieves all enabled modules for a page
func (s *moduleService) GetEnabledModulesByPageID(pageID uint) ([]models.Module, error) {
	modules, err := s.moduleRepo.GetEnabledModulesByPageID(pageID)
	if err != nil {
		return nil, fmt.Errorf("failed to get enabled modules by page ID: %w", err)
	}
	return modules, nil
}
