package models

import (
	"time"

	"gorm.io/datatypes"
)

type Template struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	TemplateName        string         `gorm:"type:varchar(255);not null" json:"template_name"`
	TemplateDescription *string        `gorm:"type:text" json:"template_description,omitempty"`
	IsOfficial          bool           `gorm:"not null;default:false" json:"is_official"`
	IsPublic            bool           `gorm:"not null;default:false" json:"is_public"`
	CreatorUserID       *uint          `gorm:"index" json:"creator_user_id,omitempty"`
	PreviewImageURL     *string        `gorm:"type:varchar(255)" json:"preview_image_url,omitempty"`
	TemplateData        datatypes.JSON `gorm:"type:jsonb;not null" json:"template_data"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	Creator *User `gorm:"foreignKey:CreatorUserID;references:ID" json:"creator,omitempty"`
}

// TableName specifies the table name for GORM
func (Template) TableName() string {
	return "templates"
}
