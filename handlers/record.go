package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFullName dummy type to return query params
func GetFullName(ctx *gin.Context) {
	record := struct {
		First    string
		Last     string
		FullName string
	}{
		ctx.Param("first"),
		ctx.Param("last"),
		ctx.Param("first") + " " + ctx.Param("last"),
	}

	ctx.JSON(http.StatusOK, record)
}
