package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maximdanilchenko/coins/wallet"
)

type (
	CreateAccountRequest struct {
		Account Account `json:"account"`
	}
	GetAccountsResponse struct {
		Accounts []Account `json:"accounts"`
	}
)

func decodeCreateAccountReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, wallet.NewErrHttp(err.Error(), http.StatusBadRequest)
	}
	// Request data validation. TODO: move it to another middleware or use some validation instruments aka swagger etc
	if req.Account.Balance.IsNegative() {
		return nil, wallet.NewErrHttp("amount should be positive", http.StatusBadRequest)
	}
	if req.Account.Balance.Exponent() != -2 {
		return nil, wallet.NewErrHttp("should be 2 decimal places", http.StatusBadRequest)
	}
	if req.Account.Currency == "" {
		return nil, wallet.NewErrHttp("currency should not be empty", http.StatusBadRequest)
	}
	if req.Account.Name == "" {
		return nil, wallet.NewErrHttp("name should not be empty", http.StatusBadRequest)
	}
	return req, nil
}

func decodeGetAccountsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}