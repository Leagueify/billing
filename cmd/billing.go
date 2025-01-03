package main

import (
	"github.com/leagueify/billing/internal/config"
	"github.com/leagueify/billing/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewServer(cfg).Start()
}
