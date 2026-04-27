package security

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const mockSecret = "my_secret_key_test-mock"

func generateValidToken(role string) string {
	claims := jwt.MapClaims{
		"sub":  "1234567890",
		"role": role,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(mockSecret))
	return "Bearer " + tokenString
}

func TestGenerateToken_DadosValidos_DeveRetornarTokenString(t *testing.T) {
	// Arrange
	userID := "user-123"
	role := "admin"

	// Act
	tokenStr, err := GenerateToken(userID, role, mockSecret)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)
}
func TestGenerateToken_SecretVazio_DeveGerarTokenMesmoAssim(t *testing.T) {
	// Nota: O JWT permite secret vazio (embora não seguro),
	// o importante é garantir que a biblioteca não dê panic.

	// Arrange
	userID := "user-123"
	role := "viewer"
	secretVazio := ""

	// Act
	tokenStr, err := GenerateToken(userID, role, secretVazio)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)
}
