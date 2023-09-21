package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/something-to-start-with/api-server-go/internal/v1/api"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	api.SetupRoutes(router)
	return router
}

func main() {
	router := setupRouter()
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
