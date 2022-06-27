package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	// Disable Console Color
	r := gin.Default()

	// index default route
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{

		// check for API v1 availability
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		// retrieve secret from database
		v1.GET("/retrieve/:slug", load)

		// save secret to database
		v1.POST("/save", save)
	}

	return r
}
