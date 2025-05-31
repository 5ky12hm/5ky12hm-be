package router

import (
	"api/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(e *gin.Engine) *gin.Engine {
	dynamodbClient := client.NewDynamodbClient()

	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	SetupBlogRouter(e, dynamodbClient)

	return e
}
