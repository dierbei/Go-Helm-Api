package helmcontroller

import (
	"net/http"
	"strconv"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListRepository() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		pageSize := ctx.DefaultQuery("pageSize", "10")

		iPage, err := strconv.Atoi(page)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, err)
			return
		}

		iPageSize, err := strconv.Atoi(pageSize)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1002, err)
			return
		}

		res, err := h.Svc.ListRepository(iPage, iPageSize)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1002, err)
			return
		}

		response.NewResponse(ctx).Success(res)
	}
}
