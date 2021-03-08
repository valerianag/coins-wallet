package payment

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

// CreatePayment creates new payment
func (s service) CreatePayment(ctx context.Context, payment Payment) error {
	logger := log.With(s.logger, "method", "CreatePayment")

	if err := s.storage.CreatePayment(ctx, payment); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	level.Info(logger).Log()
	return nil
}

// CreatePayment returns all existing payments
func (s service) GetPayments(ctx context.Context) ([]Payment, error) {
	logger := log.With(s.logger, "method", "GetPayments")

	payments, err := s.storage.GetPayments(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	level.Info(logger).Log()
	return payments, nil
}
