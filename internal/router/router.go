package router

import (
	helmcontroller "githup.com/dierbei/go-helm-api/internal/controller/helm"
	helmservice "githup.com/dierbei/go-helm-api/internal/service/helm"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	// Helm
	helmGroup := engine.Group("/helm")
	helmService := helmservice.New()
	helmcontroller.New(helmGroup, helmService)
}
