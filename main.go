package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"secure-backend/db"
	"secure-backend/http"
)

func main() {
	log.SetPrefix("Backend API: ")
	log.Print("Starting application")
	gin.SetMode(gin.ReleaseMode)
	r := http.SetupRouter()
	ctx := context.TODO()

	db.ConnectDB()
	db.ConnectRedis(ctx)

	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())
	}

}
