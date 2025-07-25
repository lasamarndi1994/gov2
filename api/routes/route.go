package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	api.POST("/login", handler.HandleLogin)
	api.POST("/register", handler.HandleRegister)

	// auth := r.Group("/api", AuthMiddleware())
	// auth.GET("/dashboard", dashboardHandler)

	return router
}
