package service

import (
	"fmt"
	"log"
	"time"

	"link-nest/internal/models"
	"link-nest/internal/repository"
)

// PaymentService defines the interface for payment-related business logic
type PaymentService interface {
	CreatePayment(userID uint, subscriptionPlan string, amount float64, currency, gateway, transactionID string) (*models.Payment, error)
	GetPaymentsByUserID(userID uint) ([]models.Payment, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(paymentRepo repository.PaymentRepository) PaymentService {
	return &paymentService{paymentRepo: paymentRepo}
}

// CreatePayment handles payment creation
func (s *paymentService) CreatePayment(userID uint, subscriptionPlan string, amount float64, currency, gateway, transactionID string) (*models.Payment, error) {
	payment := &models.Payment{
		UserID:           userID,
		SubscriptionPlan: subscriptionPlan,
		Amount:           amount,
		Currency:         currency,
		PaymentGateway:   gateway,
		TransactionID:    transactionID,
		Status:           "pending",
		PaymentDate:      time.Now(),
	}

	// Create payment in database
	if err := s.paymentRepo.CreatePayment(payment); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	log.Printf("Payment created: %s for user %s", payment.TransactionID, payment.UserID)
	return payment, nil
}

// GetPaymentsByUserID retrieves all payments for a user
func (s *paymentService) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	payments, err := s.paymentRepo.GetPaymentsByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payments by user ID: %w", err)
	}
	return payments, nil
}
