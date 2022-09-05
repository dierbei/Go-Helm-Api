package helmcontroller

import (
	"log"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func (h *handler) PullChart() gin.HandlerFunc {
	// RepoURL, version, destDir, chartName string

	return func(ctx *gin.Context) {
		c := action.NewPullWithOpts(action.WithConfig(&action.Configuration{}))

		//if registryUsername != "" {
		//	client.Username = registryUsername
		//}
		//if registryPassword != "" {
		//	client.Password = registryPassword
		//}
		c.Settings = &cli.EnvSettings{}
		c.Version = "0.1.0"
		c.RepoURL = "http://175.24.198.168:8080"
		c.DestDir = "tmp"
		name, err := c.Run("tchart")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}
}
