package helmservice

import (
	"context"
	gohelmclient "github.com/mittwald/go-helm-client"
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
	"helm.sh/helm/v3/pkg/release"
)

func (s *service) InstallOrUpgradeRelease(namespace, repository, release, chart, version string) (*release.Release, error) {
	client, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace)
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
