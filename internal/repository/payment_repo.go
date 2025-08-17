package repository

import (
	"link-nest/internal/models"

	"gorm.io/gorm"
)

// PaymentRepository defines the interface for payment data operations
type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentByID(paymentID uint) (*models.Payment, error)
	GetPaymentsByUserID(userID uint) ([]models.Payment, error)
	GetPaymentByTransactionID(transactionID string) (*models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository creates a new PaymentRepository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// CreatePayment creates a new payment in the database
func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

// GetPaymentByID retrieves a payment by its ID
func (r *paymentRepository) GetPaymentByID(paymentID uint) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, paymentID).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

// GetPaymentsByUserID retrieves all payments for a user
func (r *paymentRepository) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.Where("user_id = ?", userID).Order("payment_date DESC").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// GetPaymentByTransactionID retrieves a payment by transaction ID
func (r *paymentRepository) GetPaymentByTransactionID(transactionID string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.Where("transaction_id = ?", transactionID).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
