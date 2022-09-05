package helmservice

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/mysql"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
)

func (s *service) SelectRepository(search *helmrepo.Repository) (*helmrepo.Repository, error) {
	var (
		db = mysql.GetDb()
		//data = helmrepo.NewRepository()
	)

	if err := db.Debug().Limit(1).Find(search).Error; err != nil {
		return nil, err
	}

	return search, nil
}
