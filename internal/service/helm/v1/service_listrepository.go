package v1

import (
	"math"

	"githup.com/dierbei/go-helm-api/internal/pkg/mysql"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
)

func (s *service) ListRepository(page, pageSize int) (*helmrepo.ListRepositoryResponse, error) {
	var (
		db     = mysql.GetDb()
		res    = helmrepo.NewListRepositoryResponse()
		data   = make(helmrepo.RepositoryList, 0)
		limit  = pageSize
		offset = (page - 1) * pageSize
		count  int64
	)

	if page == -1 || pageSize == -1 {
		offset = -1
		limit = -1
	}

	db = db.Model(data).Count(&count)
	db = db.Offset(offset).Limit(limit)
	db.Order("created_at desc").Find(&data)

	res.Data = data
	res.Total = count
	res.Page = page
	res.PageSize = pageSize
	res.TotalPage = int(math.Ceil(float64(count) / float64(pageSize)))
	return res, nil
}
