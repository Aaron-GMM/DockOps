package config

import (
	"os"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUrl     string
	JWTSecret string
}

var Logger = logger.NewLogger("Config")

func Load() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		Logger.Error("Aviso: arquivo .env não encontrado. Usando variáveis de sistema.")
	}

	return &AppConfig{
		DBUrl:     os.Getenv("DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
