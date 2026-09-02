package router

import (
	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "github.com/Aaron-GMM/DockOps/docs"
)

func InitRouter(db *gorm.DB, pub core.MessagePublisher) {
	var router *gin.Engine = gin.Default()
	initializeRouter(router, db, pub)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
