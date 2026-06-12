package security

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type opaRequest struct {
	Input opaInputData `json:"input"`
}
type opaInputData struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Role   string `json:"role"`
}

type opaResponse struct {
	Result bool `json:"result"`
}

func OPAMiddleware(opaURL string) gin.HandlerFunc {

	return func(c *gin.Context) {
		roleInterface, ok := c.Get("role")
		if !ok {
			log.ErrorF("Role not found in context. The AuthMiddleware is running?")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Permission denied: No role provided"})
			return
		}
		request := opaRequest{
			Input: opaInputData{
				Method: c.Request.Method,
				Path:   c.Request.URL.Path,
				Role:   roleInterface.(string),
			},
		}
		body, _ := json.Marshal(request)
		resp, err := http.Post(opaURL, "application/json", bytes.NewBuffer(body))
		if err != nil || resp.StatusCode != http.StatusOK {
			log.ErrorF("OPA error. Request: %v, Response: %v", request, resp)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Authorization server unavailable"})
			return
		}
		defer resp.Body.Close()
		var opaResponse opaResponse
		if err := json.NewDecoder(resp.Body).Decode(&opaResponse); err != nil {
			log.ErrorF("OPA error decode. Request: %v, Response: %v", request, resp)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Authorization server unavailable"})
			return
		}
		if !opaResponse.Result {
			log.ErrorF("OPA error acess deined to role '%s'. Request: %v, Response: %v", request.Input.Role, request, opaResponse)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden by Policy"})
			return
		}
		c.Next()
	}
}
