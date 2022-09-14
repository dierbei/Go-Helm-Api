package v1

import (
	"net/http"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListReleaseHistories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		release := ctx.Query("release")
		max := 10

		releaseHistories, err := h.Svc.ListReleaseHistories(release, namespace, max)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1001, err)
			return
		}

		response.NewResponse(ctx).Success(releaseHistories)
	}
}
