package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/gin-gonic/gin"
)

type ContainerHandler struct {
	publisher core.MessagePublisher
}

var log = logger.NewLogger("API-Handler")

func NewContainerHandler(publisher core.MessagePublisher) *ContainerHandler {
	return &ContainerHandler{
		publisher: publisher,
	}
}

func (h *ContainerHandler) CreateContainer(c *gin.Context) {
	payload := core.ContainerPayload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.WarningF("Bad rquest (JSON No formated): %v")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Formato de Json invalido",
			"details": err.Error(),
		})
		return
	}
	log.InforF("Request accepted for create Container: %s (Imagem: %s)", payload.Name, payload.Image)
	msgBytes, err := json.Marshal(payload)
	if err != nil {
		log.ErrorF("Error al serializar json: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.publisher.Publish(c.Request.Context(), "container_tasks", msgBytes)

	if err != nil {
		log.ErrorF("Error al publish in queue: 'container_task':%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Not published in queue",
			"details": err.Error(),
		})
		return
	}

	log.InforF("Request processed for create Container: %s (Imagem: %s)", payload.Name, payload.Image)
	c.JSON(http.StatusAccepted, gin.H{
		"message": string(msgBytes),
		"status":  "processing",
	})
}
