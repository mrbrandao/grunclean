package main

import (
	"grunclean/model"
)

func main() {

	model.Flags()

	//	fmt.Println("Starting ...")
	model.Actions(model.Action, model.Type)
}
