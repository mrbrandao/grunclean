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
	teste := model.ListJobs(model.Url, model.Token)
	fmt.Println(teste)

}
