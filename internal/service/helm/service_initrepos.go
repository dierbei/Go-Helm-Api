package helmservice

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
	"helm.sh/helm/v3/pkg/repo"
	"log"
)

func (s *service) InitRepos() error {
	var (
		settings = helmclient.GetHelmSettings()
	)

	repositories, err := s.ListRepository(-1, -1)
	if err != nil {
		return err
	}

	repoEntryList := make([]*repo.Entry, 0)
	helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, "default")
	for _, r := range repositories.Data {
		entry := repo.Entry{
			Name:     r.Name,
			URL:      r.URL,
			Username: r.Username,
			Password: r.Password,
		}
		repoEntryList = append(repoEntryList, &entry)

		if err := helmClient.AddOrUpdateChartRepo(entry); err != nil {
			log.Println(err)
		}
	}
	settings.InitRepos(repoEntryList)

	return nil
}
