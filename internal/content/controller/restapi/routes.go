package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/something-to-start-with/api-server-go/internal/content"
)

func SetupRoutes(r *gin.Engine, s content.Service) {
	handler := New(s)

	v1 := r.Group("/v1")
	{
		v1.GET("/contents", handler.GetContents)
		v1.POST("/contents", handler.Create)
		v1.PUT("/contents/:id", handler.Update)
		v1.GET("/contents/:id", handler.GetById)
		v1.DELETE("/contents/:id", handler.DeleteById)
	}
}
