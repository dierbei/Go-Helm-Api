package helmclient

import (
	"encoding/base64"
	gohelmclient "github.com/mittwald/go-helm-client"
	"k8s.io/client-go/rest"
)

var (
	_helmClient gohelmclient.Client
)

func GetHelmClient(apiServer, Token, CA, namespace string) (gohelmclient.Client, error) {
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
	_helmClient = _client

	return _client, nil
}
