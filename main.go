package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro/services"
)

func main(){
	//注册服务到consul
	 consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.104:8500"),
	)
	ginRouter := gin.Default()
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("GET","/prods", func(context *gin.Context) {
			//context.String(200,"user api")
			context.JSON(200,services.NewProdList(5))
		})
	}

	server := web.NewService(
		web.Name("gomicroservice"),
		//web.Address(":8081"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
		)
	server.Init()
	server.Run()
}
