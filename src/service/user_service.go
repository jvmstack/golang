package service

import (
	"gin_template/src/common/constant"
	"gin_template/src/common/response"
	"gin_template/src/domain/vo"
	"gin_template/src/repository"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	FindUserInfo(*gin.Context, *int64)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{repository}
}

func (us UserService) FindUserInfo(ctx *gin.Context, userId *int64) {
	user, _ := us.userRepository.FindUser(userId)

	if user == nil {
		response.Success(ctx, constant.DataIsNilCode, constant.DataIsNilMsg, nil)
	} else {
		userVo := &vo.UserVO{
			UserId:   user.UserId,
			UserName: user.UserName,
			Age:      user.Age,
			Gender:   user.Gender,
		}

		response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, userVo)
	}
}
