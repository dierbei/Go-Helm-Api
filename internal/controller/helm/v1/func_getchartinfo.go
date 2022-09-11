package v1

import (
	"fmt"
	"net/http"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
)

func (h *handler) GetChartInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("chart")
		if len(name) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, fmt.Errorf("chart name can not be empty"))
			return
		}

		// all, readme, values, chart
		info := ctx.Query("info")
		if len(info) == 0 {
			info = string(action.ShowAll)
		}

		version := ctx.Query("version")

		chartInfo, err := h.Svc.GetChartInfo(name, info, version)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1002, err)
			return
		}

		response.NewResponse(ctx).Success(chartInfo)
	}
}
