package main

import (
	"github.com/Aaron-GMM/DockOps/internal/api/router"
	"github.com/Aaron-GMM/DockOps/internal/config"
	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/storage/postgres"
)

func main() {
	Logger := logger.NewLogger("DockOps")
	Logger.Info("DockOps is starting...")

	cfg := config.Load()

	db, err := postgres.Connect(cfg.DBUrl)
	if err != nil {
		Logger.ErrorF("DockOps failed to connect: %v", err)
	}
	router.InitRouter(db)

}
