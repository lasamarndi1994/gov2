package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/handler"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Ping test
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/api/login", handler.HandleLogin)
	// auth := r.Group("/api", AuthMiddleware())
	// auth.GET("/dashboard", dashboardHandler)

	return router
}
