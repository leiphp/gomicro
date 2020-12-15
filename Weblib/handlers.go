package Weblib

import (
	"context"
	"github.com/gin-gonic/gin"
	"gomicro/Services"
)

func GetProdsList(ginCtx *gin.Context)  {
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500,gin.H{"status":err.Error()})
	}else {
		prodRes,_ := prodService.GetProdsList(context.Background(),&prodReq)
		ginCtx.JSON(200,gin.H{"data":prodRes.Data})
	}
}