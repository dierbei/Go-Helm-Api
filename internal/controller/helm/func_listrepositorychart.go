package helmcontroller

import (
	"net/http"

	"githup.com/dierbei/go-helm-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListRepositoryChart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repo := ctx.Param("repository")
		version := ctx.Query("version")   // chart version
		versions := ctx.Query("versions") // if "true", all versions
		keyword := ctx.Query("keyword")   // search keyword

		// default stable
		if version == "" {
			version = ">0.0.0"
		}

		chartList, err := h.Svc.ListRepoChart(repo, version, keyword, versions)
		if err != nil {
			response.NewResponse(ctx).Error(
				http.StatusInternalServerError,
				1001,
				err,
			)
			return
		}

		response.NewResponse(ctx).Success(chartList)
	}
}
