package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/billing/handler"
	"github.com/leagueify/billing/internal/config"
	"github.com/leagueify/billing/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	billingRouter := handler.BillingRouter()

	router.Handle(
		"/billing/", http.StripPrefix(
			"/billing", billingRouter,
		),
	)

	// define server config
	billingServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := billingServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
