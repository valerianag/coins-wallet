package account

import (
	"context"
	"encoding/json"
	"net/http"
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
		return nil, err
	}
	return req, nil
}

func decodeGetAccountsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}