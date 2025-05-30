package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(e *gin.Engine) *gin.Engine {
	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	e = SetupBlogRouter(e)
	return e
}
