package helmrepo

type repoChartElement struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	AppVersion  string `json:"app_version"`
	Description string `json:"description"`
}

func NewRepoChartElement() *repoChartElement {
	return &repoChartElement{}
}

type RepoChartList []*repoChartElement

type ListRepoChartResponse struct {
	Total    int           `json:"total"`
	RepoName string        `json:"repo_name"`
	Data     RepoChartList `json:"data"`
}

func NewListRepoChartResponse() *ListRepoChartResponse {
	return &ListRepoChartResponse{}
}

type AddRepositoryRequest struct {
	// 仓库名称
	Name string `json:"name" binding:"required"`
	// 仓库地址
	URL string `json:"url" binding:"required"`
	// 仓库用户名
	Username string `json:"username"`
	// 仓库密码
	Password string `json:"password"`
}

func NewAddRepositoryRequest() *AddRepositoryRequest {
	return &AddRepositoryRequest{}
}

type UpdateRepositoryRequest struct {
	// 主键
	ID int `json:"id" binding:"required"`
	// 仓库名称
	Name string `json:"name" binding:"required"`
	// 仓库地址
	URL string `json:"url" binding:"required"`
	// 仓库用户名
	Username string `json:"username"`
	// 仓库密码
	Password string `json:"password"`
}

func NewUpdateRepositoryRequest() *UpdateRepositoryRequest {
	return &UpdateRepositoryRequest{}
}

type UpdateRepositoryResponse struct {
	RowsAffected int64 `json:"rows_affected"`
}

func NewUpdateRepositoryResponse() *UpdateRepositoryResponse {
	return &UpdateRepositoryResponse{}
}

type ListRepositoryResponse struct {
	// 总记录数
	Total int64 `json:"total"`
	// 总页数
	TotalPage int `json:"total_page"`
	// 页码
	Page int `json:"page"`
	// 每页条数
	PageSize int `json:"page_size"`
	// 数据
	Data RepositoryList `json:"data"`
}

func NewListRepositoryResponse() *ListRepositoryResponse {
	return &ListRepositoryResponse{}
}
