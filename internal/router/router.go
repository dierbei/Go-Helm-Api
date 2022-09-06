package router

import (
	helmcontroller "githup.com/dierbei/go-helm-api/internal/controller/helm"
	helmservice "githup.com/dierbei/go-helm-api/internal/service/helm"
	"log"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	// Helm
	helmService := helmservice.New()
	if err := helmService.InitRepos(); err != nil {
		log.Fatal(err)
	}

	helmGroup := engine.Group("/helm")
	helmcontroller.New(helmGroup, helmService)
}
