package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

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

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
