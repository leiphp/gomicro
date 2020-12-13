package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro/helpers"
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
		v1Group.Handle("POST","/prods", func(context *gin.Context) {
			//context.String(200,"user api")
			var pr helpers.ProdsRequest
			err := context.Bind(&pr)
			if err != nil || pr.Size <= 0 {
				pr = helpers.ProdsRequest{Size:2}
			}
			context.JSON(200,
				gin.H{"data":services.NewProdList(pr.Size)})
		})
	}

	server := web.NewService(
		web.Name("gomicroservice"),
		//web.Address(":8081"),
		web.Handler(ginRouter),//路由
		web.Metadata(map[string]string{"protocol": "http"}), // 添加这行代码
		web.Registry(consulReg),
		)
	server.Init()
	server.Run()
}
