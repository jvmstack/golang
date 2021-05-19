package repository

import (
	"gin_template/src/common/database"
	"gin_template/src/domain/entity"
)

type IUserRepository interface {
	FindUser(*int64) (*entity.User, error)
}

type UserManagerRepository struct {
}

func NewUserManagerRepository() IUserRepository {
	return &UserManagerRepository{}
}

func (u UserManagerRepository) FindUser(userId *int64) (*entity.User, error) {
	var user entity.User
	database.DB.Where("user_id = ?", &userId).First(&user)

	return &user, nil
}
