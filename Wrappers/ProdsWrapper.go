package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"gomicro/Services"
	"strconv"
)

type ProdsWrapper struct {
	client.Client
}

func(this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	cmdName := req.Service()+"."+req.Endpoint()
	configA := hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  0,
		RequestVolumeThreshold: 0,
		SleepWindow:            0,
		ErrorPercentThreshold:  0,
	}
	hystrix.ConfigureCommand(cmdName,configA)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx,req,rsp)
	}, func(e error) error {
		defaultProds(rsp)
		return nil
	})
}

func NewProdsWrapper (c client.Client) client.Client {
	return &ProdsWrapper{c}
}

//测试方法
func newProd(id int32,pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id,ProdName:pname}
}

func defaultProds(rsp interface{}) {
	models := make([]*Services.ProdModel,0)
	var i int32
	for i = 0; i < 5; i ++ {
		models = append(models,newProd(20+i,"prodname"+strconv.Itoa(100+int(i))))
	}
	result := rsp.(*Services.ProdListResponse)
	result.Data=models
}
