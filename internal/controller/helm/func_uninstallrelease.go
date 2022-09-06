package helmcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"
)

func (h *handler) UninstallRelease() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		release := ctx.Param("release")
		namespace := ctx.Param("namespace")

		if err := h.Svc.UninstallRelease(namespace, release); err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1001, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	}
}
