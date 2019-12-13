package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetQuotes route function
func GetQuotes(ctx *gin.Context) {
	response := struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}{"allochi", 45}

	ctx.JSON(http.StatusOK, response)
}

// GetTenders route function
func GetTenders(ctx *gin.Context) {
	ctx.String(http.StatusOK, "These are tenders!")
}

// GetAccounts route function
func GetAccounts(ctx *gin.Context) {
	users, err := ioutil.ReadFile("data/users.json")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.String(http.StatusOK, strings.ReplaceAll(string(users), `\`, ``))
}

// GetReportSources route function
func GetReportSources(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.String(http.StatusOK, "The report_id: %d", id)
}
