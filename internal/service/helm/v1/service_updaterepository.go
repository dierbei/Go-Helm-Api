package v1

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
	"githup.com/dierbei/go-helm-api/internal/pkg/mysql"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
	"helm.sh/helm/v3/pkg/repo"
	"log"
)

func (s *service) UpdateRepository(req *helmrepo.UpdateRepositoryRequest) (*helmrepo.UpdateRepositoryResponse, error) {
	var (
		db       = mysql.GetDb()
		data     = helmrepo.NewRepository()
		settings = helmclient.GetHelmSettings()
		search   = helmrepo.NewRepository()
	)

	search.ID = req.ID
	if _, err := s.SelectRepository(search); err != nil {
		return nil, err
	}

	data.ID = req.ID
	data.Name = req.Name
	data.URL = req.URL
	data.Username = req.Username
	data.Password = req.Password

	result := db.Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}

	go func() {
		go func() {
			settings.UpdateRepo(&repo.Entry{
				Name:     data.Name,
				URL:      data.URL,
				Username: data.Username,
				Password: data.Password,
			})

			helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, "default")
			if err != nil {
				log.Println(err)
				return
			}
			helmClient.AddOrUpdateChartRepo(repo.Entry{
				Name:     data.Name,
				URL:      data.URL,
				Username: data.Username,
				Password: data.Password,
			})
		}()
	}()

	response := helmrepo.NewUpdateRepositoryResponse()
	response.RowsAffected = result.RowsAffected
	return response, nil
}
