package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	token = "sometoken"
	url   = "somerundeckurl"
)

func main() {
	fmt.Println("Starting ...")
	response, err := http.Get(url + "/projects?authtoken=" + token)
	if err != nil {
		fmt.Printf("Failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	fmt.Println("Staring json...")
	response, _ := http.NewRequest("POST", url+"/tokens/isb?authtoken="+token, nil)

	//	jsonData := map[string]string{"user": "teste"}
	//	jsonValue, _ := json.Marshal(jsonData)
	//	response, err = http.Post(url+"/tokens?authtoken="+token, "application/json", bytes.NewBuffer(jsonValue))
	//	if err != nil {
	//		fmt.Printf("Failed with error %s\n", err)
	//	} else {
	//		data, _ := ioutil.ReadAll(response.Body)
	//		fmt.Println(string(data))
	//	}

}
