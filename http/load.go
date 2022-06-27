package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func load(c *gin.Context) {

	slug := c.Param("slug")
	log.Println("SLUG: " + slug)

	c.String(http.StatusOK, "ok")

}
