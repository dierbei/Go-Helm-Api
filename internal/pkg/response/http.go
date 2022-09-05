package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *responseData {
	return &responseData{ctx: ctx}
}

func (r *responseData) Success(data interface{}) {
	response := gin.H{
		"code": http.StatusOK,
		"data": data,
	}
	r.ctx.JSON(http.StatusOK, response)
}

func (r *responseData) Error(code, debugCode int, err error) {
	response := gin.H{
		"code":       code,
		"debug_code": debugCode,
		"msg":        err.Error(),
	}
	r.ctx.JSON(code, response)
}
