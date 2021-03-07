package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateAccount endpoint.Endpoint
	GetAccounts   endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateAccount: makeCreateAccountEndpoint(s),
		GetAccounts:   makeGetAccountsEndpoint(s),
	}
}

func makeCreateAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		err := s.CreateAccount(ctx, req.Account)
		return nil, err
	}
}

func makeGetAccountsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		accounts, err := s.GetAccounts(ctx)
		if err != nil {
			return nil, err
		}
		return GetAccountsResponse{Accounts: accounts}, nil
	}
}
