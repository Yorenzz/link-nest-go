package models

import (
	"time"
)

type UserPage struct {
	ID                   uint    `gorm:"primaryKey" json:"id"`
	UserID               uint    `gorm:"not null;index" json:"user_id"`
	PageSlug             string  `gorm:"type:varchar(100);unique;not null" json:"page_slug"`
	PageTitle            string  `gorm:"type:varchar(255);not null;default:'我的主页'" json:"page_title"`
	PageDescription      *string `gorm:"type:text" json:"page_description,omitempty"`
	SEOKeywords          *string `gorm:"type:text" json:"seo_keywords,omitempty"`
	SocialShareImageURL  *string `gorm:"type:varchar(255)" json:"social_share_image_url,omitempty"`
	CustomDomain         *string `gorm:"type:varchar(255);unique" json:"custom_domain,omitempty"`
	IsActive             bool    `gorm:"not null;default:true" json:"is_active"`
	TemplateID           *uint   `gorm:"index" json:"template_id,omitempty"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	User     User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Template *Template `gorm:"foreignKey:TemplateID;references:ID" json:"template,omitempty"`
}

// TableName specifies the table name for GORM
func (UserPage) TableName() string {
	return "user_pages"
}
