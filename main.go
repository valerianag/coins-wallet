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

	"github.com/maximdanilchenko/coins/wallet"
	"github.com/maximdanilchenko/coins/wallet/account"
	"github.com/maximdanilchenko/coins/wallet/payment"
)

// main function to configure and start server
func main() {
	// Create logger
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
	config, err := wallet.NewAppConfig()
	if err != nil {
		level.Error(logger).Log("config error", err)
		return
	}

	// Create db connections
	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", config.PostgresDSN)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	// Handle syscalls
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	ctx := context.Background()

	// Creat needed handlers from all services
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

	accountsEndpoints := account.MakeEndpoints(accountsService)
	paymentsEndpoints := payment.MakeEndpoints(paymentsService)

	accountsHandler := account.MakeHttpHandlers(ctx, accountsEndpoints)
	paymentsHandler := payment.MakeHttpHandlers(ctx, paymentsEndpoints)

	// Run server
	go func() {
		fmt.Println("listening on port", config.AppPort)
		sm := http.NewServeMux()
		sm.Handle("/accounts", accountsHandler)
		sm.Handle("/payments", paymentsHandler)
		errs <- http.ListenAndServe(config.AppPort, sm)
	}()

	level.Error(logger).Log("exit", <-errs)
}
