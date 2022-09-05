package helmcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// localhost:8080/helm/release/:name/namespace/:namespace?kube_context
func (h *handler) UninstallRelease() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//name := ctx.Param("release")
		//namespace := ctx.Param("namespace")
		//kubeContext := ctx.Query("kube_context")
		//kubeConfig := ctx.Query("kube_config")

		// todo 修改settings
		//settings.KubeToken = testClusterToken
		//settings.KubeCaFile = testClusterCA
		//settings.KubeAPIServer = testClusterApiServer

		//actionConfig, err := actionConfigInit(namespace)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusInternalServerError, gin.H{})
		//	return
		//}

		//clientGo, err := clientgo.GetClientGo()
		//if err != nil {
		//	return
		//}
		//client := action.NewUninstall(actionConfig)
		//_, err = client.Run(name)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusInternalServerError, gin.H{})
		//	return
		//}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	}
}
