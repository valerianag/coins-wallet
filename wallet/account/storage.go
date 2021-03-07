package account

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
)

type storage struct {
	db     *sql.DB
	logger log.Logger
}

func NewStorage(db *sql.DB, logger log.Logger) Storage {
	return &storage{
		db:     db,
		logger: log.With(logger, "storage", "sql"),
	}
}

func (s storage) CreateAccount(ctx context.Context, account Account) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO accounts (name, balance, currency)
		VALUES ($1, $2, $3)`, account.Name, account.Balance.String(), account.Currency)

	return err
}

func (s storage) GetAccounts(ctx context.Context) ([]Account, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT name, balance, currency FROM accounts`)
	if err != nil {
		return nil, err
	}

	var result []Account

	defer rows.Close()
	for rows.Next() {
		var account Account

		err = rows.Scan(&account.Name, &account.Balance, &account.Currency)
		if err != nil {
			return nil, err
		}

		result = append(result, account)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
