package Weblib

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"gomicro/Services"
	"strconv"
)
//测试方法
func newProd(id int32,pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id,ProdName:pname}
}

func defaultProds() (*Services.ProdListResponse,error) {
	models := make([]*Services.ProdModel,0)
	var i int32
	for i = 0; i < 5; i ++ {
		models = append(models,newProd(20+i,"prodname"+strconv.Itoa(100+int(i))))
	}
	res:=&Services.ProdListResponse{}
	res.Data=models
	return res,nil
}

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
		}, func(e error) error {
			prodRes,err = defaultProds()
			return err
		})
		if err != nil {
			ginCtx.JSON(500,gin.H{"status":err.Error()})
		}else{
			ginCtx.JSON(200,gin.H{"data":prodRes.Data})
		}

	}
}