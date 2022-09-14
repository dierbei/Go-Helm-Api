package v1

import (
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"

	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
	"helm.sh/helm/v3/cmd/helm/search"
)

func (s *service) ListRepositoryChart(repoName, version, keyword, versions string) (*helmrepo.ListRepoChartResponse, error) {
	settings := helmclient.GetHelmSettings()

	index, err := settings.BuildSearchIndex(repoName, version)
	if err != nil {
		return nil, err
	}

	indexRes := index.All()
	if len(keyword) != 0 {
		indexRes, err = index.Search(keyword, helmclient.SearchMaxScore, false)
		if err != nil {
			return nil, err
		}
	}

	search.SortScore(indexRes)
	var versionsB bool
	if versions == "true" {
		versionsB = true
	}

	data, err := settings.ApplyConstraint(version, versionsB, indexRes)
	if err != nil {
		return nil, err
	}

	chartList := make(helmrepo.RepoChartList, 0, len(data))
	for _, v := range data {
		element := helmrepo.NewRepoChartElement()
		element.Name = v.Name
		element.Version = v.Chart.Version
		element.AppVersion = v.Chart.AppVersion
		element.Description = v.Chart.Description
		chartList = append(chartList, element)
	}

	response := helmrepo.NewListRepoChartResponse()
	response.Total = len(chartList)
	response.RepoName = repoName
	response.Data = chartList

	return response, nil
}
