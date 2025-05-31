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
	blogs, err := s.blogsRepo.FindAll()
	if err != nil {
		return nil, err
	}
	blogOverviews := []*dto.BlogOverview{}
	for _, blog := range blogs {
		blogOverviews = append(blogOverviews, &dto.BlogOverview{
			Id:        blog.Id,
			Title:     blog.Title,
			Tag:       blog.Tag,
			UpdatedAt: blog.UpdatedAt,
		})

	}
	return &dto.GetBlogsRes{
		BlogOverviews: blogOverviews,
	}, nil
}

func (s *blogService) GetBlog(req *dto.GetBlogReq) (*dto.GetBlogRes, error) {
	blog, err := s.blogsRepo.FindById(req.BlogId)
	if err != nil {
		return nil, err
	}
	// TODO: occure bad request
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
