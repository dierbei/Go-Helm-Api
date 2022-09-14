package v1

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"

	gohelmclient "github.com/mittwald/go-helm-client"
)

func (s *service) RollBack(release, namespace, repository, version string) error {
	helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace)
	if err != nil {
		return err
	}

	chartSpec := gohelmclient.ChartSpec{
		ReleaseName: release,
		ChartName:   repository + "/" + release,
		Namespace:   namespace,
		Version:     version,
	}

	return helmClient.RollbackRelease(&chartSpec)
}
