package service

import (
	"github.com/wa1kman999/cowman/pkg/user/dao"
	"github.com/wa1kman999/cowman/pkg/user/model"
)

type DomainUser interface {
	// FindOne 查询一个
	FindOne(user *model.User) (*model.User, error)
}

// DomainUserService 用户领域服务
type DomainUserService struct {
}

func NewDomainUserService() DomainUser {
	return new(DomainUserService)
}

// FindOne 查询一个
func (domain *DomainUserService) FindOne(user *model.User) (*model.User, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return nil, err
	}
	return entity.FindOne(user)
}
