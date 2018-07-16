package main

import (
	"grunclean/model"
)

func main() {

	model.Flags()

	//	fmt.Println("Starting ...")
	model.Actions(model.Action, model.Type)

	//listid := model.ListExecutions(model.Url, model.Token)
	//	model.DeleteExecution(model.ListExecutions(model.Url, model.Token))
	//	for idx, i := range listid.Executions {
	//		fmt.Printf("Indice[%d] ExecID: %d e mais algo = %+v\r\n", idx, listid.Executions[idx].Id, i)
	//fmt.Println(i)
	//	}
	//	for idx, i := range listid {
	//		fmt.Println("indice[%d] execID:[%d]", idx, i)
	//	}

}
