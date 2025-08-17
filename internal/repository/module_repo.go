package repository

import (
	"link-nest/internal/models"

	"gorm.io/gorm"
)

// ModuleRepository defines the interface for module data operations
type ModuleRepository interface {
	CreateModule(module *models.Module) error
	GetModuleByID(moduleID uint) (*models.Module, error)
	GetModulesByPageID(pageID uint) ([]models.Module, error)
	GetEnabledModulesByPageID(pageID uint) ([]models.Module, error)
}

type moduleRepository struct {
	db *gorm.DB
}

// NewModuleRepository creates a new ModuleRepository
func NewModuleRepository(db *gorm.DB) ModuleRepository {
	return &moduleRepository{db: db}
}

// CreateModule creates a new module in the database
func (r *moduleRepository) CreateModule(module *models.Module) error {
	return r.db.Create(module).Error
}

// GetModuleByID retrieves a module by its ID
func (r *moduleRepository) GetModuleByID(moduleID uint) (*models.Module, error) {
	var module models.Module
	if err := r.db.First(&module, moduleID).Error; err != nil {
		return nil, err
	}
	return &module, nil
}

// GetModulesByPageID retrieves all modules for a page
func (r *moduleRepository) GetModulesByPageID(pageID uint) ([]models.Module, error) {
	var modules []models.Module
	if err := r.db.Where("page_id = ?", pageID).Order("module_order").Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

// GetEnabledModulesByPageID retrieves all enabled modules for a page
func (r *moduleRepository) GetEnabledModulesByPageID(pageID uint) ([]models.Module, error) {
	var modules []models.Module
	if err := r.db.Where("page_id = ? AND is_enabled = ?", pageID, true).Order("module_order").Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}
