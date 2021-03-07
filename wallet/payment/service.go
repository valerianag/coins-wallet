package payment

import (
	"context"
)

type Service interface {
	CreatePayment(ctx context.Context, payment Payment) error
	GetPayments(ctx context.Context) ([]Payment, error)
}
