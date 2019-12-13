package handlers

import "github.com/gin-gonic/gin"

// Authenticate middleware
func Authenticate() gin.HandlerFunc {
	return gin.BasicAuth(
		gin.Accounts{
			"foo":     "bar",
			"allochi": "bar",
		},
	)
}
