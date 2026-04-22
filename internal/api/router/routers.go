package router

import (
	"github.com/Aaron-GMM/DockOps/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initializeRouter(router *gin.Engine, db *gorm.DB) {
	_ = postgres.NewEventRepository(db)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/create")

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong - DockOps API is alive!",
			})
		})
	}
}
