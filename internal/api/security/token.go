package security

import (
	"time"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/golang-jwt/jwt/v5"
)

var log = logger.NewLogger("security-token")

func GenerateToken(userID string, role string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		"iat":  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinamos o token com a nossa chave secreta
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.ErrorF("Erro ao assinar o token JWT para o usuario %s: %v", userID, err)
		return "", err
	}

	log.InforF("Token gerado com sucesso para o ID: %s", userID)
	return tokenString, nil
}
