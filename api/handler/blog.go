package handler

import (
	"api/dto"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler interface {
	GetBlogs(*gin.Context)
	GetBlog(*gin.Context)
}

type blogHandler struct {
	blogService service.BlogService
}

func NewBlogHandler(
	blogService service.BlogService,
) BlogHandler {
	return &blogHandler{
		blogService: blogService,
	}
}

func (h *blogHandler) GetBlogs(c *gin.Context) {
	res, err := h.blogService.GetBlogs()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *blogHandler) GetBlog(c *gin.Context) {
	req := &dto.GetBlogReq{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	res, err := h.blogService.GetBlog(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
