package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"githup.com/dierbei/go-helm-api/config"
	"githup.com/dierbei/go-helm-api/internal/pkg/response"
	"io/ioutil"
	"net/http"
)

func (h *handler) DeleteChart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		chart := ctx.Query("chart")
		if len(chart) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1001, fmt.Errorf("invalid chart name"))
			return
		}

		version := ctx.Query("version")
		if len(version) == 0 {
			response.NewResponse(ctx).Error(http.StatusBadRequest, 1002, fmt.Errorf("invalid chart version"))
			return
		}

		_cfg := config.GetConfig().Chartmuseum
		client := http.Client{}
		request, err := http.NewRequest(http.MethodDelete, _cfg.Address+"/api/charts/"+chart+"/"+version, nil)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1003, err)
			return
		}

		res, err := client.Do(request)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1004, err)
			return
		}
		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			response.NewResponse(ctx).Error(http.StatusInternalServerError, 1005, err)
			return
		}

		response.NewResponse(ctx).Success(string(b))
	}
}
