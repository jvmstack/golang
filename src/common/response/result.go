package response

import (
	"gin_template/src/common/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultVO struct {
	Code    constant.ResponseCode `json:"code"`
	Msg     constant.ResponseMsg  `json:"msg"`
	Success bool                  `json:"success"`
	Data    interface{}           `json:"data"`
}

/**
 * 请求成功函数
 */
func Success(ctx *gin.Context, code constant.ResponseCode, msg constant.ResponseMsg, data interface{}) {
	resp := &ResultVO{Code: code, Msg: msg, Success: true, Data: data}
	ctx.JSON(http.StatusOK, resp)
}

/**
 * 请求失败函数
 */
func Failure(ctx *gin.Context, code constant.ResponseCode, msg constant.ResponseMsg, data interface{}) {
	resp := &ResultVO{Code: code, Msg: msg, Success: false, Data: data}
	ctx.JSON(http.StatusInternalServerError, resp)
}
