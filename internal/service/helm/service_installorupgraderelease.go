package helmservice

import (
	"context"
	gohelmclient "github.com/mittwald/go-helm-client"
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

func (s *service) InstallOrUpgradeRelease(namespace, repository, release, chart, version string) (*release.Release, error) {
	var (
		search = helmrepo.NewRepository()
	)

	search.Name = repository
	_repository, err := s.SelectRepository(search)
	if err != nil {
		return nil, err
	}

	client, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace, repo.Entry{
		Name:     _repository.Name,
		URL:      _repository.URL,
		Username: _repository.Username,
		Password: _repository.Password,
	})
	if err != nil {
		return nil, err
	}

	chartSpec := &gohelmclient.ChartSpec{
		ReleaseName: release,
		ChartName:   repository + "/" + chart,
		Namespace:   namespace,
		Version:     version,
	}

	return client.InstallOrUpgradeChart(context.Background(), chartSpec, nil)
}
