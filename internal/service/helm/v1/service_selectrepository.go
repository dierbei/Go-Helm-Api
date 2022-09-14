package v1

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/mysql"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
)

func (s *service) SelectRepository(search *helmrepo.Repository) (*helmrepo.Repository, error) {
	var (
		db = mysql.GetDb()
	)

	if err := db.Debug().Limit(1).Find(search).Error; err != nil {
		return nil, err
	}

	return search, nil
}
