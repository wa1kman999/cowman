package user

import (
	"github.com/wa1kman999/cowman/pkg/user/model"
	"github.com/wa1kman999/cowman/pkg/user/service"
)

type AppService struct{}

func NewAppFormService() *AppService {
	return &AppService{}
}

// Login 登陆
func (app *AppService) Login(param *model.User) (*model.User, error) {
	userService := service.NewDomainUserService()
	user, err := userService.FindOne(param)
	if err != nil {
		return nil, err
	}
	return user, nil
}
