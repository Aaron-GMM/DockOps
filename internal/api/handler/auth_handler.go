package handler

import (
	"net/http"

	"github.com/Aaron-GMM/DockOps/internal/api/security"
	"github.com/Aaron-GMM/DockOps/internal/config"
	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo  core.UserRepository
	jwtSecret string
}

func NewAuthHandler(userRepo core.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		jwtSecret: config.Load().JWTSecret,
	}
}

// LoginRequest model for authentication
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"aarongm"`
	Password string `json:"password" binding:"required" example:"secret123"`
}

// Login godoc
// @Summary Authenticate user
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username e Password são obrigatórios"})
		return
	}

	user, err := h.userRepo.GetByUsername(c.Request.Context(), req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if !security.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha incorreta"})
		return
	}

	token, err := security.GenerateToken(user.ID, user.Role, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  user.Role,
	})
}
