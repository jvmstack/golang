package controller

import (
	"gin_template/src/common/constant"
	"gin_template/src/common/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func UserRegister(userGrp *gin.RouterGroup) {
	userController := &UserController{}
	userGrp.Use().POST("/findUser", userController.findUser)
}

func (c UserController) findUser(ctx *gin.Context) {

	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, gin.H{"userName": "月影"})
}
