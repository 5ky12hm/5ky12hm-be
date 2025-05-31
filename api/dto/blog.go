package dto

type GetBlogsRes struct {
	BlogOverviews []*BlogOverview `json:"blogs"`
}

type GetBlogReq struct {
	BlogId string `uri:"blogId"`
}

type GetBlogRes struct {
	BlogOverview
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

type BlogOverview struct {
	Id        string `json:"blogId"`
	Title     string `json:"title"`
	Tag       string `json:"tag"`
	UpdatedAt string `json:"updatedAt"`
}
