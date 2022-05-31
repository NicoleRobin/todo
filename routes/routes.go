package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/nicolerobin/todo/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.Cors(), gin.Recovery())
	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("/user/register", nil)
	}
	return r
}
