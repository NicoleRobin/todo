package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nicolerobin/todo/api"
	"github.com/nicolerobin/todo/docs"
	"github.com/nicolerobin/todo/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// pprof
	pprof.Register(r)

	// cookie based session
	store := cookie.NewStore([]byte("cookie-secret"))
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.Cors(), gin.Recovery())

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ug := r.Group("api/v1")
	{
		// 用户操作
		ug.GET("/hello", api.Hello)
		ug.POST("user/register", api.UserRegister)
		ug.POST("user/login", api.UserLogin)
		authed := ug.Group("/")
		authed.Use(middleware.JWT())
		{
			/*
				// 任务操作
				authed.GET("tasks", api.ListTasks)
				authed.POST("task", api.CreateTask)
				authed.GET("task:id", api.ShowTask)
				authed.DELETE("task/:id", api.DeleteTask)
				authed.PUT("task/:id", api.UpdateTask)
				authed.POST("search", api.SearchTasks)
			*/
		}
	}
	return r
}
