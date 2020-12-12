package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main(){
	//注册服务到consul
	 consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.104:8500"),
	)
	ginRouter := gin.Default()
	ginRouter.Handle("GET","/user", func(context *gin.Context) {
		context.String(200,"user api")
	})
	ginRouter.Handle("GET","/news", func(context *gin.Context) {
		context.String(200,"news api")
	})
	server := web.NewService(
		web.Name("gomicroservice"),
		web.Address(":8081"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
		)

	server.Run()
}
