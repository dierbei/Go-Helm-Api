package helmservice

import (
	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
)

func (s *service) UninstallRelease(namespace, release string) error {
	helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace)
	if err != nil {
		return err
	}

	return helmClient.UninstallReleaseByName(release)
}
