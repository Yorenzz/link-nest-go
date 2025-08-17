package repository

import (
	"link-nest/internal/models"

	"gorm.io/gorm"
)

// TemplateRepository defines the interface for template data operations
type TemplateRepository interface {
	CreateTemplate(template *models.Template) error
	GetTemplateByID(templateID uint) (*models.Template, error)
	GetOfficialTemplates() ([]models.Template, error)
	GetPublicTemplates() ([]models.Template, error)
}

type templateRepository struct {
	db *gorm.DB
}

// NewTemplateRepository creates a new TemplateRepository
func NewTemplateRepository(db *gorm.DB) TemplateRepository {
	return &templateRepository{db: db}
}

// CreateTemplate creates a new template in the database
func (r *templateRepository) CreateTemplate(template *models.Template) error {
	return r.db.Create(template).Error
}

// GetTemplateByID retrieves a template by its ID
func (r *templateRepository) GetTemplateByID(templateID uint) (*models.Template, error) {
	var template models.Template
	if err := r.db.First(&template, templateID).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// GetOfficialTemplates retrieves all official templates
func (r *templateRepository) GetOfficialTemplates() ([]models.Template, error) {
	var templates []models.Template
	if err := r.db.Where("is_official = ?", true).Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

// GetPublicTemplates retrieves all public templates
func (r *templateRepository) GetPublicTemplates() ([]models.Template, error) {
	var templates []models.Template
	if err := r.db.Where("is_public = ?", true).Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}
