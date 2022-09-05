package helmcontroller

import (
	"github.com/gin-gonic/gin"
	"githup.com/dierbei/go-helm-api/internal/pkg/response"
	"net/http"
)

func (h *handler) InstallOrUpgradeRelease() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		repository := ctx.Query("repository")
		release := ctx.Query("release")
		chart := ctx.Query("chart")
		version := ctx.Query("version")

		res, err := h.Svc.InstallOrUpgradeRelease(namespace, repository, release, chart, version)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1001, err)
			return
		}

		response.NewResponse(ctx).Success(res)
	}
}
