package router

import (
	"os"

	"github.com/Aaron-GMM/DockOps/internal/api/handler"
	"github.com/Aaron-GMM/DockOps/internal/api/security"
	"github.com/Aaron-GMM/DockOps/internal/config"
	"github.com/Aaron-GMM/DockOps/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initializeRouter(router *gin.Engine, db *gorm.DB) {
	repo := postgres.NewEventRepository(db)

	jwtSecret := config.Load().JWTSecret
	han := handler.NewContainerHandler(nil, repo)
	// A URL do OPA. "dockops.authz" é o package do rego, e "allow" é a variável que criamos.
	// Busca da variável de ambiente, mas se estiver vazia (rodando no terminal), usa localhost
	opaURL := os.Getenv("OPA_URL")
	if opaURL == "" {
		opaURL = "http://localhost:8181/v1/data/dockops/authz/allow"
	}

	v1 := router.Group("/api/v1")
	{
		containers := v1.Group("/containers")

		// APLICANDO A SEGURANÇA NA ROTA DE CRIAÇÃO!
		// 1. AuthMiddleware valida o token e extrai a role.
		// 2. OPAMiddleware pergunta ao OPA se essa role pode fazer POST.
		containers.POST("/",
			security.AuthMiddleware(jwtSecret),
			security.OPAMiddleware(opaURL),
			han.CreateContainer,
		)
	}
}
