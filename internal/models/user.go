package models

import (
	"time"
)

type User struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	Username          string     `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email             string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash      string     `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber       *string    `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	AuthProvider      string     `gorm:"type:varchar(50);not null;default:'email_password'" json:"auth_provider"`
	IsProUser         bool       `gorm:"not null;default:false" json:"is_pro_user"`
	LastLoginAt       *time.Time `gorm:"type:timestamptz" json:"last_login_at,omitempty"`
	ProfilePictureURL *string    `gorm:"type:varchar(255)" json:"profile_picture_url,omitempty"`
	PreferredLanguage string     `gorm:"type:varchar(10);not null;default:'en-US'" json:"preferred_language"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}
