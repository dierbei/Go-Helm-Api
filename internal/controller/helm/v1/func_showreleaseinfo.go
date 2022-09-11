package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"githup.com/dierbei/go-helm-api/internal/pkg/response"
	"net/http"
)

func (h *handler) ShowReleaseInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("release")
		if len(name) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, fmt.Errorf("invalid release name"))
			return
		}

		namespace := ctx.Param("namespace")
		if len(namespace) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1002, fmt.Errorf("invalid release namespace"))
			return
		}

		info := ctx.Query("info")
		if info == "" {
			info = "values"
		}
		output := ctx.Query("output")

		data, err := h.Svc.ShowReleaseInfo(name, namespace, info, output)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1003, err)
			return
		}

		response.NewResponse(ctx).Success(data)
	}
}
