package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type content struct {
	ID   int
	Body string
}

var contents = []content{
	{ID: 1, Body: "content1"},
	{ID: 2, Body: "content2"},
}

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/contents", getContents)
	}
}

func getContents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contents)
}
