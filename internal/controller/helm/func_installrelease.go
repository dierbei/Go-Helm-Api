package helmcontroller

import (
	"github.com/pkg/errors"
	"net/http"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) InstallOrUpgradeRelease() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		repository := ctx.Query("repository")
		release := ctx.Query("release")
		chart := ctx.Query("chart")
		version := ctx.Query("version")

		if len(namespace) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, errors.New("invalid namespace"))
			return
		}

		if len(repository) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1002, errors.New("invalid repository"))
			return
		}

		if len(release) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1003, errors.New("invalid release"))
			return
		}

		if len(chart) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1004, errors.New("invalid chart"))
			return
		}

		if len(version) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1005, errors.New("invalid version"))
			return
		}

		res, err := h.Svc.InstallOrUpgradeRelease(namespace, repository, release, chart, version)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1006, err)
			return
		}

		response.NewResponse(ctx).Success(res)
	}
}
