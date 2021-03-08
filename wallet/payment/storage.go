package payment

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/shopspring/decimal"
)

type storage struct {
	db     *sql.DB
	logger log.Logger
}

type account struct {
	id       int
	balance  decimal.Decimal
	currency string
}

var (
	errNotEqualCurrencies = errors.New("not equal currencies")
	errNotEnoughBalance   = errors.New("not enough balance")
)

func NewStorage(db *sql.DB, logger log.Logger) Storage {
	return &storage{
		db:     db,
		logger: log.With(logger, "storage", "sql"),
	}
}

func (s storage) CreatePayment(ctx context.Context, payment Payment) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err = createPaymentTx(tx, ctx, payment); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func createPaymentTx(tx *sql.Tx, ctx context.Context, payment Payment) error {
	accountFrom, err := queryAccountForUpdate(ctx, tx, payment.FromAccount)
	if err != nil {
		return err
	}
	accountTo, err := queryAccountForUpdate(ctx, tx, payment.ToAccount)
	if err != nil {
		return err
	}

	if accountFrom.currency != accountTo.currency {
		return errNotEqualCurrencies
	}
	if accountFrom.balance.LessThan(payment.Amount) {
		return errNotEnoughBalance
	}

	_, err = tx.ExecContext(ctx,
		`UPDATE accounts SET balance = $1 WHERE id = $2`,
		accountFrom.balance.Sub(payment.Amount).String(), accountFrom.id,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`UPDATE accounts SET balance = $1 WHERE id = $2`,
		accountTo.balance.Add(payment.Amount).String(), accountTo.id,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO payments (account_from, account_to, amount)
		VALUES ($1, $2, $3)`, accountFrom.id, accountTo.id, payment.Amount,
	)
	if err != nil {
		return err
	}

	return nil
}

func queryAccountForUpdate(ctx context.Context, tx *sql.Tx, accountName string) (account account, err error) {
	row := tx.QueryRowContext(ctx, `
		SELECT id, balance, currency FROM accounts WHERE name = $1 FOR UPDATE`, accountName)
	err = row.Scan(&account.id, &account.balance, &account.currency)

	return
}

func (s storage) GetPayments(ctx context.Context) ([]Payment, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT af.name, at.name, amount FROM payments p
		JOIN accounts af on p.account_from = af.id
		JOIN accounts at on p.account_to = at.id`)
	if err != nil {
		return nil, err
	}

	var result []Payment

	defer rows.Close()
	for rows.Next() {
		var payment Payment

		err = rows.Scan(&payment.FromAccount, &payment.ToAccount, &payment.Amount)
		if err != nil {
			return nil, err
		}

		result = append(result, payment)
	}

	return result, nil
}
