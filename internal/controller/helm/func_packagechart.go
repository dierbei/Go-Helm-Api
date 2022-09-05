package helmcontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
)

func (h handler) PackageChart() gin.HandlerFunc {
	destDir := "tmp/mychart"

	return func(ctx *gin.Context) {
		if _, err := action.NewPackage().Run(destDir, nil); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	}
}
