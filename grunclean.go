package main

import (
	"grunclean/model"
)

func main() {

	model.Flags()

	//	fmt.Println("Starting ...")
	//version := model.Version(Url)
	//fmt.Printf("Using Api Version: %+v on url %s\r\n", version, Url)
	//	model.ListJobs(model.Url, model.Token)
	//	fmt.Println(teste)

	//	list := model.ListExecutions(model.Url, model.Token)
	//	for i := 0; i < len(list.Executions); i++ {
	//		fmt.Println(list.Executions[i].Id)
	//	}
	//model.Test()
	model.Actions(model.Action, model.Type)

}
