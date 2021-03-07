package account

import (
	"context"
)

type Service interface {
	CreateAccount(ctx context.Context, account Account) error
	GetAccounts(ctx context.Context) ([]Account, error)
}
