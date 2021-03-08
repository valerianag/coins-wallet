package payment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maximdanilchenko/coins/wallet"
)

type (
	CreatePaymentRequest struct {
		Payment Payment `json:"payment"`
	}
	GetPaymentsResponse struct {
		Payments []Payment `json:"payments"`
	}
)

func decodeCreatePaymentReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreatePaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, wallet.NewErrHttp(err.Error(), http.StatusBadRequest)
	}
	// Request data validation. TODO: move it to another middleware or use some validation instruments aka swagger etc
	if !req.Payment.Amount.IsPositive() {
		return nil, wallet.NewErrHttp("amount should be positive", http.StatusBadRequest)
	}
	if req.Payment.Amount.Exponent() != -2 {
		return nil, wallet.NewErrHttp("should be 2 decimal places", http.StatusBadRequest)
	}
	if req.Payment.FromAccount == "" {
		return nil, wallet.NewErrHttp("from_account should not be empty", http.StatusBadRequest)
	}
	if req.Payment.ToAccount == "" {
		return nil, wallet.NewErrHttp("to_account should not be empty", http.StatusBadRequest)
	}
	return req, nil
}

func decodeGetPaymentsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
