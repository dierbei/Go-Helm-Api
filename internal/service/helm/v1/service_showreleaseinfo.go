package v1

import (
	"fmt"

	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"
)

func (s *service) ShowReleaseInfo(name, namespace, info string) (interface{}, error) {
	if info != "all" && info != "values" {
		return nil, fmt.Errorf("bad info %s, release info only support all/values", info)
	}

	helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace)
	if err != nil {
		return nil, err
	}

	if info == "values" {
		values, err := helmClient.GetReleaseValues(name, true)
		if err != nil {
			return nil, err
		}
		return values, nil
	}

	release, err := helmClient.GetRelease(name)
	if err != nil {
		return nil, err
	}

	return release, nil
}
