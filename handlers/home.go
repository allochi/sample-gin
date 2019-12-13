package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home route function
func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title": "gin-gonic sample",
	})
}
