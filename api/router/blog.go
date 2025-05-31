package router

import (
	"api/client"
	"api/handler"
	"api/repository"
	"api/service"

	"github.com/gin-gonic/gin"
)

func SetupBlogRouter(e *gin.Engine, dynamodbClient client.DynamodbClient) {
	blogsGroup := e.Group("/blogs")
	{
		h := handler.NewBlogHandler(
			service.NewBlogService(
				repository.NewBlogsRepository(dynamodbClient),
			),
		)
		blogsGroup.GET("", h.GetBlogs)
		blogsGroup.GET("/:blogId", h.GetBlog)
	}
}
