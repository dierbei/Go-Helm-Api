package helmclient

import (
	"encoding/base64"
	gohelmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/client-go/rest"
)

func GetHelmClient(apiServer, Token, CA, namespace string, entry repo.Entry) (gohelmclient.Client, error) {
	decodeToken, err := base64.StdEncoding.DecodeString(Token)
	if err != nil {
		return nil, err
	}
	decodeCa, err := base64.StdEncoding.DecodeString(CA)
	if err != nil {
		return nil, err
	}

	opt := &gohelmclient.RestConfClientOptions{
		Options: &gohelmclient.Options{
			Namespace:        namespace,
			RepositoryCache:  "tmp/.helmcache",
			RepositoryConfig: "tmp/.helmrepo",
			Debug:            true,
			Linting:          true,
			DebugLog: func(format string, v ...interface{}) {
			},
		},
		RestConfig: &rest.Config{
			Host:        apiServer,
			BearerToken: string(decodeToken),
			TLSClientConfig: rest.TLSClientConfig{
				CAData: decodeCa,
			},
		},
	}

	_client, err := gohelmclient.NewClientFromRestConf(opt)
	if err != nil {
		return nil, err
	}

	if err := _client.AddOrUpdateChartRepo(entry); err != nil {
		return nil, err
	}

	return _client, nil
}
