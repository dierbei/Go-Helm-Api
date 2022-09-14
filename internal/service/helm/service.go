package helmservice

import (
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
	v1 "githup.com/dierbei/go-helm-api/internal/service/helm/v1"

	"helm.sh/helm/v3/pkg/release"
)

type Service interface {
	ListRepositoryChart(repoName, version, keyword, versions string) (*helmrepo.ListRepoChartResponse, error)

	AddRepository(req *helmrepo.AddRepositoryRequest) error

	UpdateRepository(req *helmrepo.UpdateRepositoryRequest) (*helmrepo.UpdateRepositoryResponse, error)

	SelectRepository(search *helmrepo.Repository) (*helmrepo.Repository, error)

	InstallOrUpgradeRelease(namespace, repository, release, chart, version string) (*release.Release, error)

	ListRepository(page, pageSize int) (*helmrepo.ListRepositoryResponse, error)

	UninstallRelease(namespace, release string) error

	GetChartInfo(chart, info, version string) (interface{}, error)

	InitRepos() error

	ShowReleaseInfo(release, namespace, info string) (interface{}, error)

	ListReleaseHistories(release, namespace string, max int) (interface{}, error)

	RollBack(release, namespace, repository, version string) error
}

func New() Service {
	return v1.NewV1Service()
}
