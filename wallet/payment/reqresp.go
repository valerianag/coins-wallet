package payment

import (
	"context"
	"encoding/json"
	"net/http"
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
		return nil, err
	}
	return req, nil
}

func decodeGetPaymentsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
