package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/maximdanilchenko/coins/wallet/account"
	"github.com/maximdanilchenko/coins/wallet/payment"
)

const dbsource = "postgresql://postgres:postgres@localhost:5432/wallet_db?sslmode=disable"

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "wallet",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	ctx := context.Background()

	var (
		accountsService account.Service
		paymentsService payment.Service
	)
	{
		accountsStorage := account.NewStorage(db, logger)
		paymentsStorage := payment.NewStorage(db, logger)

		accountsService = account.NewService(accountsStorage, logger)
		paymentsService = payment.NewService(paymentsStorage, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	accountsEndpoints := account.MakeEndpoints(accountsService)
	paymentsEndpoints := payment.MakeEndpoints(paymentsService)

	accountsHandler := account.NewHTTPServer(ctx, accountsEndpoints)
	paymentsHandler := payment.NewHTTPServer(ctx, paymentsEndpoints)

	go func() {
		fmt.Println("listening on port", ":8080")
		sm := http.NewServeMux()
		sm.Handle("/accounts", accountsHandler)
		sm.Handle("/payments", paymentsHandler)
		errs <- http.ListenAndServe(":8080", sm)
	}()

	level.Error(logger).Log("exit", <-errs)
}
