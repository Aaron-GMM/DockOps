package main

import (
	"github.com/Aaron-GMM/DockOps/internal/api/router"
	"github.com/Aaron-GMM/DockOps/internal/config"
	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/messaging/rabbitmq"
	"github.com/Aaron-GMM/DockOps/internal/storage/postgres"
)

// @title           DockOps API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.linkedin.com/in/aaron-moreira
// @contact.email  aarongmdev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	Logger := logger.NewLogger("DockOps")
	Logger.Info("DockOps is starting...")

	cfg := config.Load()

	// 1. Conexão com Banco de Dados
	db, err := postgres.Connect(cfg.DBUrl)
	if err != nil {
		Logger.ErrorF("DockOps failed to connect to DB: %v", err)
	}

	// 2. Conexão com RabbitMQ
	rabbitConn, err := rabbitmq.NewConnection(cfg.RabbitMQUrl)
	var pub *rabbitmq.RabbitPublisher
	if err != nil {
		Logger.ErrorF("DockOps failed to connect to RabbitMQ: %v", err)
	} else {
		defer rabbitConn.Close()
		pub = rabbitmq.NewRabbitPublisher(*rabbitConn)
	}

	// 3. Inicialização do Router com Injeção de Dependências
	router.InitRouter(db, pub)
}
