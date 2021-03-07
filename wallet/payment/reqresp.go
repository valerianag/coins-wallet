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
	return req, nil
}

func decodeGetPaymentsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
