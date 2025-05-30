package router

import (
	"api/handler"
	"api/service"

	"github.com/gin-gonic/gin"
)

func SetupBlogRouter(e *gin.Engine) *gin.Engine {
	blogsGroup := e.Group("/blogs")
	{
		h := handler.NewBlogHandler(
			service.NewBlogService(),
		)
		blogsGroup.GET("", h.GetBlogs)
		blogsGroup.GET("/:blogId", h.GetBlog)
	}
	return e
}
