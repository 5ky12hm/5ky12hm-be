package service

import (
	"api/dto"
)

type BlogService interface {
	GetBlogs() (*dto.GetBlogsRes, error)
	GetBlog(*dto.GetBlogReq) (*dto.GetBlogRes, error)
}

type blogService struct {
}

func NewBlogService() BlogService {
	return &blogService{}
}

func (s *blogService) GetBlogs() (*dto.GetBlogsRes, error) {
	return &dto.GetBlogsRes{
		Blogs: []*dto.BlogOverview{
			{
				Id:    "1",
				Title: "testTitle1",
			},
			{
				Id:    "2",
				Title: "testTitle2",
			},
		},
	}, nil
}

func (s *blogService) GetBlog(req *dto.GetBlogReq) (*dto.GetBlogRes, error) {
	return &dto.GetBlogRes{
		BlogOverview: dto.BlogOverview{
			Id:    "1",
			Title: "testTitle",
		},
		Body: "testBody",
	}, nil
}
