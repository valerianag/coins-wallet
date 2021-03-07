package account

import (
	"context"

	"github.com/shopspring/decimal"
)

type Account struct {
	Name     string          `json:"name"`
	Balance  decimal.Decimal `json:"balance"`
	Currency string          `json:"currency"`
}

type Storage interface {
	CreateAccount(ctx context.Context, account Account) error
	GetAccounts(ctx context.Context) ([]Account, error)
}
