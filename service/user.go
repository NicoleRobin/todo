package service

import (
	"github.com/nicolerobin/todo/model"
	"github.com/nicolerobin/todo/pkg/e"
	"github.com/nicolerobin/todo/serializer"
)

type UserService struct {
	UserName string `from:"user_name" json:"user_name" binding:"required,min=3,max=15" exmaple:"FanOne"`
	Password string `from:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (us *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", us.UserName).First(&user).Count(&count)
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 设置密码
	user.UserName = us.UserName
	if err := user.SetPassword(us.Password); err != nil {
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
