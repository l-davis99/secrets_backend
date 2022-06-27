package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"secure-backend/db"
)

type IncomingSave struct {
	Msg string
}

func save(c *gin.Context) {

	// Match request to struct set above
	var req IncomingSave
	ctx := context.TODO()

	// set to JSON
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	// Generate a slug for user to obtain the message
	slug := uuid.New().String()[:8]
	c.JSON(200, slug)

	db.RedisClient.Set(ctx, slug, req.Msg, 0)

}
