package postgres

import (
	"github.com/Aaron-GMM/DockOps/internal/api/security"
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

	err = db.AutoMigrate(&EventModel{}, &UserModel{})
	if err != nil {
		log.ErrorF("Erro ao executar AutoMigrate: %v", err)
		return nil, err
	}

	seedUsers(db)

	log.InfoF("Conexão com PostgreSQL estabelecida com sucesso!")
	return db, nil
}

func seedUsers(db *gorm.DB) {
	adminHash, _ := security.HashPassword("admin123")
	devHash, _ := security.HashPassword("dev123")
	viewerHash, _ := security.HashPassword("viewer123")

	users := []UserModel{
		{ID: "admin-01", Username: "admin", Password: adminHash, Role: "admin"},
		{ID: "dev-01", Username: "dev", Password: devHash, Role: "developer"},
		{ID: "viewer-01", Username: "viewer", Password: viewerHash, Role: "viewer"},
	}

	for _, u := range users {
		var count int64
		db.Model(&UserModel{}).Where("username = ?", u.Username).Count(&count)
		if count == 0 {
			db.Create(&u)
		}
	}
}
