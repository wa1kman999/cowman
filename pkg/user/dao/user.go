package dao

import (
	"github.com/wa1kman999/cowman/global"
	"github.com/wa1kman999/cowman/pkg/user/model"
	"gorm.io/gorm"
)

type User interface {
	// FindOne 查询一个
	FindOne(u *model.User) (*model.User, error)
}

type UserEntity struct {
	dao *gorm.DB
}

func NewUserEntity() (User, error) {
	return &UserEntity{
		dao: global.CMMysql,
	}, nil
}

// FindOne 通过名字查询
func (entity *UserEntity) FindOne(u *model.User) (*model.User, error) {
	var user model.User
	err := entity.dao.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
