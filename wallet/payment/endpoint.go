package payment

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/maximdanilchenko/coins/wallet"
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
		if err != nil {
			return nil, wallet.NewErrHttp(err.Error(), http.StatusBadRequest)
		}
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
