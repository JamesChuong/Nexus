package auth_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	AuthToken string `header:"Authorization" binding:"required"`
	RequestID string `header:"X-Request-ID"`
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestHeader HeaderParams

		if err := c.ShouldBindHeader(&requestHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required headers missing"})
		}
		token := requestHeader.AuthToken

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		}
		
	}
}
