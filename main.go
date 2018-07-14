package main

import (
	"fmt"
	"grunclean/model"
)

func main() {

	model.Flags()

	fmt.Println("Starting ...")
	//version := model.Version(Url)
	//fmt.Printf("Using Api Version: %+v on url %s\r\n", version, Url)
	model.ListJobs(model.Url, model.Token)
	//	fmt.Println(teste)
	fmt.Println("LUIZA")
	listOlder := model.ListOlderExecutions(model.Url, model.Token)
	for i := 0; i < len(listOlder.Executions); i++ {
		fmt.Println(listOlder.Executions[i].Id)
	}
	//model.Test()

}
