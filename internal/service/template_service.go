package service

import (
	"fmt"
	"log"

	"link-nest/internal/models"
	"link-nest/internal/repository"

	"gorm.io/datatypes"
)

// TemplateService defines the interface for template-related business logic
type TemplateService interface {
	CreateTemplate(name, description string, isOfficial, isPublic bool, creatorUserID *uint, templateData datatypes.JSON) (*models.Template, error)
	GetOfficialTemplates() ([]models.Template, error)
	GetPublicTemplates() ([]models.Template, error)
}

type templateService struct {
	templateRepo repository.TemplateRepository
}

// NewTemplateService creates a new TemplateService
func NewTemplateService(templateRepo repository.TemplateRepository) TemplateService {
	return &templateService{templateRepo: templateRepo}
}

// CreateTemplate handles template creation
func (s *templateService) CreateTemplate(name, description string, isOfficial, isPublic bool, creatorUserID *uint, templateData datatypes.JSON) (*models.Template, error) {
	template := &models.Template{
		TemplateName:    name,
		IsOfficial:      isOfficial,
		IsPublic:        isPublic,
		CreatorUserID:   creatorUserID,
		TemplateData:    templateData,
	}

	if description != "" {
		template.TemplateDescription = &description
	}

	// Create template in database
	if err := s.templateRepo.CreateTemplate(template); err != nil {
		return nil, fmt.Errorf("failed to create template: %w", err)
	}

	log.Printf("Template created: %s", template.TemplateName)
	return template, nil
}

// GetOfficialTemplates retrieves all official templates
func (s *templateService) GetOfficialTemplates() ([]models.Template, error) {
	templates, err := s.templateRepo.GetOfficialTemplates()
	if err != nil {
		return nil, fmt.Errorf("failed to get official templates: %w", err)
	}
	return templates, nil
}

// GetPublicTemplates retrieves all public templates
func (s *templateService) GetPublicTemplates() ([]models.Template, error) {
	templates, err := s.templateRepo.GetPublicTemplates()
	if err != nil {
		return nil, fmt.Errorf("failed to get public templates: %w", err)
	}
	return templates, nil
}
