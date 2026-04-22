package postgres

import (
	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logger.NewLogger("DataBase")

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.ErrorF("Erro ao conectar no banco de dados: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&EventModel{})
	if err != nil {
		log.ErrorF("Erro ao executar AutoMigrate: %v", err)
		return nil, err
	}

	log.ErrorF("Conexão com PostgreSQL estabelecida com sucesso!")
	return db, nil
}
