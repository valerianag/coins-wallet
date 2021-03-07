package payment

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreatePayment endpoint.Endpoint
	GetPayments   endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreatePayment: makeCreatePaymentEndpoint(s),
		GetPayments:   makeGetPaymentsEndpoint(s),
	}
}

func makeCreatePaymentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePaymentRequest)
		err := s.CreatePayment(ctx, req.Payment)
		return nil, err
	}
}

func makeGetPaymentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		Payments, err := s.GetPayments(ctx)
		if err != nil {
			return nil, err
		}
		return GetPaymentsResponse{Payments: Payments}, nil
	}
}
