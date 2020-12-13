package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro/datamodels"
	"io/ioutil"
	"log"
	"net/http"
	myhttp "github.com/micro/go-plugins/client/http"
)

//
func callAPI2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("gomicroservice","/v1/prods",
		datamodels.ProdsRequest{Size:3})
	var rsp datamodels.ProdListResponse
	err := myClient.Call(context.Background(),req,&rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.GetData())
}

//原始调用方式
func callAPI(addr string, path string, method string) (string, error) {
	req,_ := http.NewRequest(method,"http://"+addr+path,nil)
	client := http.DefaultClient
	res,err:=client.Do(req)
	if err != nil {
		return "",err
	}
	defer res.Body.Close()
	buf,_ := ioutil.ReadAll(res.Body)
	return string(buf),nil

}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.104:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
		)
	callAPI2(mySelector)
	//for {
	//	getService,err := consulReg.GetService("gomicroservice")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	next := selector.RoundRobin(getService)
	//	node, _ := next()
	//	callRes,err := callAPI(node.Address,"/v1/prods","GET")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(callRes)
	//}

}