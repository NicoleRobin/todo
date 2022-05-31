package main

import (
	"github.com/nicolerobin/todo/conf"
	"github.com/nicolerobin/todo/routes"

	_ "net/http/pprof"
)

func main() {
	conf.LoadConf()
	r := routes.NewRouter()
	r.GET("/hello", api.Hello)
	_ = r.Run(conf.Config.Web.Address)
}
