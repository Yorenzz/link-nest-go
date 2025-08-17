package models

import (
	"time"

	"gorm.io/datatypes"
)

type Module struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PageID        uint           `gorm:"not null;index" json:"page_id"`
	ModuleType    string         `gorm:"type:varchar(50);not null" json:"module_type"`
	ModuleTitle   *string        `gorm:"type:varchar(255)" json:"module_title,omitempty"`
	ModuleIconURL *string        `gorm:"type:varchar(255)" json:"module_icon_url,omitempty"`
	ModuleOrder   int            `gorm:"not null" json:"module_order"`
	IsEnabled     bool           `gorm:"not null;default:true" json:"is_enabled"`
	Config        datatypes.JSON `gorm:"type:jsonb;not null" json:"config"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	Page UserPage `gorm:"foreignKey:PageID;references:ID" json:"page"`
}

// TableName specifies the table name for GORM
func (Module) TableName() string {
	return "modules"
}
