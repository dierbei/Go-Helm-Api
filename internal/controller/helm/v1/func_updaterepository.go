package v1

import (
	"net/http"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateRepository() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := helmrepo.NewUpdateRepositoryRequest()
		if err := ctx.ShouldBindJSON(req); err != nil {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, err)
			return
		}

		res, err := h.Svc.UpdateRepository(req)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1002, err)
			return
		}

		response.NewResponse(ctx).Success(res)
	}
}
