package service

import (
	"api/dto"
	"api/repository"
)

type BlogService interface {
	GetBlogs() (*dto.GetBlogsRes, error)
	GetBlog(*dto.GetBlogReq) (*dto.GetBlogRes, error)
}

type blogService struct {
	blogsRepo repository.BlogsRepository
}

func NewBlogService(
	blogsRepo repository.BlogsRepository,
) BlogService {
	return &blogService{
		blogsRepo: blogsRepo,
	}
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
	blog, err := s.blogsRepo.FindById(req.BlogId)
	if err != nil {
		return nil, err
	}
	// TODO
	if blog == nil {
		return nil, nil
	}
	return &dto.GetBlogRes{
		BlogOverview: dto.BlogOverview{
			Id:        blog.Id,
			Title:     blog.Title,
			Tag:       blog.Tag,
			UpdatedAt: blog.UpdatedAt,
		},
		Body:      blog.Body,
		CreatedAt: blog.CreatedAt,
	}, nil
}
