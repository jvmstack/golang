package controller

import (
	"gin_template/src/domain/query"
	"gin_template/src/repository"
	"gin_template/src/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func UserRegister(userGrp *gin.RouterGroup) {
	userRepository := repository.NewUserManagerRepository()
	userService := service.NewUserService(userRepository)
	userController := &UserController{userService: userService}

	userGrp.Use().POST("/findUser", userController.findUser)
}

func (uc UserController) findUser(ctx *gin.Context) {
	var userQuery query.FindUserQuery
	_ = ctx.Bind(&userQuery)
	uc.userService.FindUserInfo(ctx, &userQuery.UserId)
}
