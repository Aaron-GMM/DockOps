package main

import (
	"fmt"
	"testing"

	"github.com/Aaron-GMM/DockOps/internal/api/security"
	"github.com/Aaron-GMM/DockOps/internal/config"
)

func TestGenerateTokensForManualTesting(t *testing.T) {
	ks := config.Load()

	// Usamos o '_' para ignorar o erro, já que é só um script auxiliar
	adminToken, _ := security.GenerateToken("dev12", "developer", ks.JWTSecret)
	fmt.Println("\n==============================================")
	fmt.Println("🟢 TOKEN DEVELOPER (DEVE CRIAR O CONTAINER)")
	fmt.Println("Bearer", adminToken)
	fmt.Println("==============================================")
	fmt.Println() // Pula uma linha
	viewerToken, _ := security.GenerateToken("vi11", "viewer", ks.JWTSecret)
	fmt.Println("==============================================")
	fmt.Println("🔴 TOKEN VIEWER (DEVE SER BLOQUEADO PELO OPA)")
	fmt.Println("Bearer", viewerToken)
	fmt.Println("==============================================")
}
