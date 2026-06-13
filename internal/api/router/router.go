package router

import (
	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, pub core.MessagePublisher) {
	var router *gin.Engine = gin.Default()
	initializeRouter(router, db, pub)

	router.Run(":8080")
}
