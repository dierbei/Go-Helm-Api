package v1

import "githup.com/dierbei/go-helm-api/internal/pkg/helmclient"

func (s *service) ListReleaseHistories(release, namespace string, max int) (interface{}, error) {
	helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace)
	if err != nil {
		return nil, err
	}

	return helmClient.ListReleaseHistory(release, max)
}
