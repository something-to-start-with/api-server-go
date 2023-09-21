package main

import (
	"github.com/gin-gonic/gin"
	"github.com/something-to-start-with/api-server-go/internal/handler"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	handler.Routes(router)

	return router
}

func main() {
	router := setupRouter()
	err := router.Run()
	if err != nil {
		return
	}
}
