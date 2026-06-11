package security

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupMockRouter(secret string) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/rota-protegida", AuthMiddleware(secret), func(c *gin.Context) {
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")

		c.JSON(http.StatusOK, gin.H{
			"message": "Acesso Permitido",
			"userID":  userID,
			"role":    role,
		})
	})
	return router
}
func TestAuthMiddleware_TokenValido_DevePermitirAcessoAteHandler(t *testing.T) {
	// Arrange
	router := setupMockRouter(mockSecret)

	// Usamos a nossa função já testada para gerar um token real
	tokenStr, _ := GenerateToken("dev-456", "developer", mockSecret)

	req, _ := http.NewRequest(http.MethodGet, "/rota-protegida", nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)         // 200 OK
	assert.Contains(t, w.Body.String(), "dev-456") // Garante que o contexto foi injetado
	assert.Contains(t, w.Body.String(), "developer")
}

func TestAuthMiddleware_SemHeaderAuthorization_DeveRetornarUnauthorized(t *testing.T) {
	// Arrange
	router := setupMockRouter(mockSecret)

	req, _ := http.NewRequest(http.MethodGet, "/rota-protegida", nil)
	// Não setamos o Header de propósito
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code) // 401 Unauthorized
	assert.Contains(t, w.Body.String(), "authorization header is missing")
}

func TestAuthMiddleware_TokenCorrompido_DeveRetornarUnauthorized(t *testing.T) {
	// Arrange
	router := setupMockRouter(mockSecret)

	req, _ := http.NewRequest(http.MethodGet, "/rota-protegida", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1Ni.corrompido.lixo")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code) // 401 Unauthorized
	assert.Contains(t, w.Body.String(), "Unauthorized")
}
