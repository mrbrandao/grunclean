package main

import (
	"fmt"
	"grunclean/model"
)

func main() {

	model.Flags()

	fmt.Println("Starting ...")
	model.Actions(model.Action, model.Type)
	//	teste := model.ListJobs(model.Url, model.Token)
	//	fmt.Println(teste)
}
