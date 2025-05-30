package dto

type GetBlogsRes struct {
	Blogs []*BlogOverview `json:"blogs"`
}

type GetBlogReq struct {
	BlogId string `url:"blogId"`
}

type GetBlogRes struct {
	BlogOverview
	Body string `json:"body"`
}

type BlogOverview struct {
	Id        string `json:"blogId"`
	Title     string `json:"title"`
	Tag       string `json:"tag"`
	UpdatedAt string `json:"updatedAt"`
}
