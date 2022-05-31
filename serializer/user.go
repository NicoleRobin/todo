package serializer

import (
	"github.com/nicolerobin/todo/model"
)

type User struct {
	Id       uint   `json:"id" form:"id" exmaple:"1"`
	UserName string `json:"user_name" form:"user_name" example:"FanOne"`
	Status   string `josn:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildUser(user model.User) User {
	return User{
		Id:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
