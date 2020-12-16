package Weblib

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"gomicro/Services"
)

//gin的方法部分
func GetProdsList(ginCtx *gin.Context)  {
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500,gin.H{"status":err.Error()})
	}else {
		//熔断代码改造
		//第一步：配置config
		configA := hystrix.CommandConfig{
			Timeout:                1000,
			MaxConcurrentRequests:  0,
			RequestVolumeThreshold: 0,
			SleepWindow:            0,
			ErrorPercentThreshold:  0,
		}
		//第二步：配置command
		hystrix.ConfigureCommand("getprods",configA)
		//第三步：执行使用Do方法，同步
		var prodRes *Services.ProdListResponse
		err := hystrix.Do("getprods",func() error {
			prodRes,err = prodService.GetProdsList(context.Background(),&prodReq)
			return err
		},nil)
		if err != nil {
			ginCtx.JSON(500,gin.H{"status":err.Error()})
		}else{
			ginCtx.JSON(200,gin.H{"data":prodRes.Data})
		}

	}
}