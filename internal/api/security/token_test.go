package security

import (
	"net/http"
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
func generateHackedToken() string {
	claims := jwt.MapClaims{"sub": "hacker-123", "role": "admin"}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims) // Algoritmo Inesperado!
	tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
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

func TestParseToken_TokenValido_DeveRetornarClaims(t *testing.T) {
	// Arrange
	header := http.Header{}
	header.Set("Authorization", generateValidToken("admin"))

	// Act
	claims, err := ParseToken(header, mockSecret)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, "1234567890", claims["sub"])
	assert.Equal(t, "admin", claims["role"])
}

func TestParseToken_HeaderVazio_DeveRetornarErro(t *testing.T) {
	// Arrange
	header := http.Header{} // Simulando requisição sem mandar o header

	// Act
	claims, err := ParseToken(header, mockSecret)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, "authorization header is missing", err.Error())
}

func TestParseToken_FormatoHeaderInvalido_DeveRetornarErro(t *testing.T) {
	// Arrange
	header := http.Header{}
	header.Set("Authorization", "Basic dXNlcjpwYXNz") // Formato errado (Basic em vez de Bearer)

	// Act
	claims, err := ParseToken(header, mockSecret)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, "invalid authorization header format", err.Error())
}

func TestParseToken_AssinaturaInvalida_DeveRetornarErro(t *testing.T) {
	// Arrange
	header := http.Header{}
	header.Set("Authorization", generateValidToken("admin"))
	secretErrado := "secret-de-outro-servidor" // Secret diferente do que gerou o token

	// Act
	claims, err := ParseToken(header, secretErrado)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "signature is invalid")
}

func TestParseToken_MetodoDeAssinaturaInesperado_DeveRetornarErro(t *testing.T) {
	// Arrange
	header := http.Header{}
	header.Set("Authorization", generateHackedToken()) // Token forjado com algoritmo "none"

	// Act
	claims, err := ParseToken(header, mockSecret)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "invalid signing method")
}
