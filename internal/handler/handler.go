package handler

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

func Routes(route *gin.Engine) {
	v1 := route.Group("/v1")
	v1.GET("/contents", getContents)
}

func getContents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contents)
}
