package models

import (
	"time"

	"gorm.io/datatypes"
)

type Payment struct {
	ID                     uint            `gorm:"primaryKey" json:"id"`
	UserID                 uint            `gorm:"not null;index" json:"user_id"`
	SubscriptionPlan       string          `gorm:"type:varchar(50);not null" json:"subscription_plan"`
	Amount                 float64         `gorm:"type:numeric(10,2);not null" json:"amount"`
	Currency               string          `gorm:"type:char(3);not null" json:"currency"`
	PaymentGateway         string          `gorm:"type:varchar(50);not null" json:"payment_gateway"`
	TransactionID          string          `gorm:"type:varchar(255);unique;not null" json:"transaction_id"`
	Status                 string          `gorm:"type:varchar(20);not null" json:"status"`
	PaymentMethodDetails   *datatypes.JSON `gorm:"type:jsonb" json:"payment_method_details,omitempty"`
	SubscriptionStartDate  *time.Time      `gorm:"type:timestamptz" json:"subscription_start_date,omitempty"`
	SubscriptionEndDate    *time.Time      `gorm:"type:timestamptz" json:"subscription_end_date,omitempty"`
	PaymentDate            time.Time       `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"payment_date"`
	CreatedAt              time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time       `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// TableName specifies the table name for GORM
func (Payment) TableName() string {
	return "payments"
}
