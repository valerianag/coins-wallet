package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/maximdanilchenko/coins/wallet"
)

func MakeHttpHandlers(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(wallet.JsonMiddleware)

	r.Methods("POST").Handler(httptransport.NewServer(
		endpoints.CreateAccount,
		decodeCreateAccountReq,
		encodeResp,
	))

	r.Methods("GET").Handler(httptransport.NewServer(
		endpoints.GetAccounts,
		decodeGetAccountsReq,
		encodeResp,
	))

	return r
}
