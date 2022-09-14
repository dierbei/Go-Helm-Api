package v1

import (
	"fmt"
	"net/http"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) RollBack() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		release := ctx.Query("release")
		repository := ctx.Query("repository")
		version := ctx.Query("version")

		if len(release) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, fmt.Errorf("invalid release name"))
			return
		}

		if len(repository) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1002, fmt.Errorf("invalid repository name"))
			return
		}

		if err := h.Svc.RollBack(release, namespace, repository, version); err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1003, err)
			return
		}

		response.NewResponse(ctx).Success(nil)
	}
}
