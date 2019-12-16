package main

import (
	"github.com/allochi/sample-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handlers.Home)
	router.LoadHTMLGlob("templates/*")

	// add `data` section to API
	// -------------------------
	data := router.Group("/data")
	data.GET("/quotes", handlers.GetQuotes)
	data.GET("/tenders", handlers.GetTenders)
	data.GET("/reports/:id/sources", handlers.GetReportSources)

	// check echo's multilevel params
	// ------------------------------
	router.GET("/name/first/:first/last/:last", handlers.GetFullName)
	// router.GET("/name/:first/:last", handlers.GetFullName)

	// add `secrets` section to API with authentication
	// ------------------------------------------------
	secrets := router.Group("/secrets")
	secrets.Use(handlers.Authenticate())
	secrets.GET("/accounts", handlers.GetAccounts)

	// Postgrest proxy
	// This is something
	router.Any("/psql/*url", handlers.PostgrestProxy)

	router.Run(":3300")
}
