package core

import (
	"encoding/hex"
	"math/rand"
)

func GenerateID() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "dockops-fallback-id"
	}
	return hex.EncodeToString(bytes)
}
