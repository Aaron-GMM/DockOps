package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Aaron-GMM/DockOps/internal/config"
	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/messaging/rabbitmq"
	"github.com/Aaron-GMM/DockOps/internal/storage/postgres"
	"github.com/Aaron-GMM/DockOps/internal/worker"
)

func main() {
	log := logger.NewLogger("Worker-Main")
	log.Info("Worker is starting...")

	cfg := config.Load()

	// 1. Conexão com Banco de Dados
	db, err := postgres.Connect(cfg.DBUrl)
	if err != nil {
		log.ErrorF("Worker failed to connect to DB: %v", err)
		os.Exit(1)
	}

	// 2. Conexão com RabbitMQ
	rabbitConn, err := rabbitmq.NewConnection(cfg.RabbitMQUrl)
	if err != nil {
		log.ErrorF("Worker failed to connect to RabbitMQ: %v", err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// 3. Iniciar o Worker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = worker.StartWorker(ctx, db, *rabbitConn)
	if err != nil {
		log.ErrorF("Worker failed to start: %v", err)
		os.Exit(1)
	}

	log.Info("Worker is running. Press Ctrl+C to stop.")

	// Esperar sinal para fechar
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info("Worker is stopping...")
}
