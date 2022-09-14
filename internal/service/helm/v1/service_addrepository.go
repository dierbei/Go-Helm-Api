package v1

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
	"githup.com/dierbei/go-helm-api/internal/pkg/mysql"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
	"helm.sh/helm/v3/pkg/repo"
	"log"
)

func (s *service) AddRepository(req *helmrepo.AddRepositoryRequest) error {
	var (
		db       = mysql.GetDb()
		data     = helmrepo.NewRepository()
		settings = helmclient.GetHelmSettings()
	)

	data.Name = req.Name
	data.URL = req.URL
	data.Username = req.Username
	data.Password = req.Password

	result := db.Create(data)
	if result.Error != nil {
		return result.Error
	}

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

	return nil
}
