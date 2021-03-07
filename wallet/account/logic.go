package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	storage Storage
	logger  log.Logger
}

func NewService(storage Storage, logger log.Logger) Service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s service) CreateAccount(ctx context.Context, account Account) error {
	logger := log.With(s.logger, "method", "CreateAccount")

	if err := s.storage.CreateAccount(ctx, account); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	level.Info(logger).Log()
	return nil
}

func (s service) GetAccounts(ctx context.Context) ([]Account, error) {
	logger := log.With(s.logger, "method", "GetAccounts")

	accounts, err := s.storage.GetAccounts(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	level.Info(logger).Log()
	return accounts, nil
}
