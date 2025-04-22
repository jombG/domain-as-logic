package domain

import (
	"errors"
	"time"
)

// Transaction представляет собой финансовую транзакцию
type Transaction struct {
	ID        string
	Amount    float64
	Currency  string
	CreatedAt time.Time
}

// Payout представляет собой выплату, которая может содержать несколько транзакций
type Payout struct {
	ID           string
	Transactions []Transaction
	TotalAmount  float64
	Currency     string
	Status       string
	CreatedAt    time.Time
	ProcessedAt  *time.Time
}

// NewPayout создает новый Payout
func NewPayout(id string, currency string) *Payout {
	return &Payout{
		ID:           id,
		Currency:     currency,
		Status:       "pending",
		CreatedAt:    time.Now(),
		Transactions: make([]Transaction, 0),
	}
}

// AddTransaction добавляет транзакцию в выплату и обновляет общую сумму
func (p *Payout) AddTransaction(transaction Transaction) error {
	// Проверяем валюту транзакции
	if transaction.Currency != p.Currency {
		return ErrInvalidCurrency
	}

	// Добавляем транзакцию
	p.Transactions = append(p.Transactions, transaction)

	// Обновляем общую сумму
	p.TotalAmount += transaction.Amount

	return nil
}

// Process обрабатывает выплату
func (p *Payout) Process() error {
	if p.Status != "pending" {
		return ErrInvalidStatus
	}

	now := time.Now()
	p.ProcessedAt = &now
	p.Status = "processed"

	return nil
}

// GetTotalAmount возвращает общую сумму выплаты
func (p *Payout) GetTotalAmount() float64 {
	return p.TotalAmount
}

// GetTransactionCount возвращает количество транзакций в выплате
func (p *Payout) GetTransactionCount() int {
	return len(p.Transactions)
}

// Ошибки домена
var (
	ErrInvalidCurrency = errors.New("invalid currency")
	ErrInvalidStatus   = errors.New("invalid status")
)
