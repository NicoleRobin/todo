package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicolerobin/todo/service"
)

// @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "user, password"
// @Success 200 {object} seiralizer.ResponseUser "{"status":200, "data":{}, "msg":"ok"}"
// @Failure 500 {object} serializer.ResponseUser "{"status": 500, "data":{}, "msg":{}, "error":"error"}"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		log.Print(err)
	}
}

// @Tags USER
// @Summary 用户登陆
// @Produce json
// @Accept json
// @Param data body service.UserService true "user, password"
// @Success 200 {object} seiralizer.ResponseUser "{"status":200, "data":{}, "msg":"ok"}"
// @Failure 500 {object} serializer.ResponseUser "{"status": 500, "data":{}, "msg":{}, "error":"error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {

}
