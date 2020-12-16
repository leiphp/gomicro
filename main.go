package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro/Services"
	"gomicro/Weblib"
)

func main(){
	//注册服务到consul
	 consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.104:8500"),
	)

	myService := micro.NewService(micro.Name("prodservice.client"))
	prodService := Services.NewProdService("prodservice",myService.Client())
	httpServer := web.NewService(
		web.Name("httpservice"),
		web.Address(":8001"),
		web.Handler(Weblib.NewGinRouter(prodService)),//路由
		//web.Metadata(map[string]string{"protocol": "http"}), // 添加这行代码
		web.Registry(consulReg),
	)
	
	httpServer.Init()
	httpServer.Run()
}
