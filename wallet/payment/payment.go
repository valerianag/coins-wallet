package payment

import (
	"context"

	"github.com/shopspring/decimal"
)

type Payment struct {
	FromAccount string          `json:"from_account"`
	ToAccount   string          `json:"to_account"`
	Amount      decimal.Decimal `json:"amount"`
}

type Storage interface {
	CreatePayment(ctx context.Context, payment Payment) error
	GetPayments(ctx context.Context) ([]Payment, error)
}
