package security

import (
	"fmt"
	"net/http"
	"strings"
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

func ParseToken(authHeader http.Header, secret string) (jwt.MapClaims, error) {
	authStr := authHeader.Get("Authorization")
	if authStr == "" {
		log.WarningF("Authorization header is missing")
		return nil, fmt.Errorf("authorization header is missing")
	}

	parts := strings.Split(authStr, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		log.WarningF("Invalid authorization header format")
		return nil, fmt.Errorf("invalid authorization header format")
	}
	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.ErrorF("Signing method is not valid: %v", token.Header["alg"])
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.WarningF("Error Parsing token: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(string); ok {
			log.Debugf("Token to ID: %s Token is Valid and Success", sub)
		} else {
			log.Debugf("Token is Valid and Success, but missing 'sub'")
		}
		return claims, nil
	}

	log.WarningF("Invalid token claims format")
	return nil, fmt.Errorf("invalid token claims format")
}
