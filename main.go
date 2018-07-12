package main

import (
	"flag"
	"fmt"
	"grunclean/model"
	"strconv"
)

var (
	Url, Token string
)

func main() {

	flag.StringVar(&Url, "url", "https://localhost:4440", "the rundeck url")
	flag.StringVar(&Token, "token", "GKrfka6yPg145IQuvvXZXbU2GxU5fKzJ", "user auth token")
	defer flag.Parse()

	fmt.Println("Starting ...")
	version := model.Version(Url)
	fmt.Printf("Using Api Version: %+v on url %s\r\n", version, Url)

	projectName := model.ListProjects(Url, Token, strconv.Itoa(version))
	fmt.Println(projectName)
	//	for _, key := range s() {
	//		strct := v.MapIndex(key)
	//		fmt.Println(key[i])
	//	}
	//}

}

//	response, err := http.Get(url + "/projects?authtoken=" + token)
//	if err != nil {
//		fmt.Printf("Failed with error %s\n", err)
//	} else {
//		data, _ := ioutil.ReadAll(response.Body)
//		fmt.Println(string(data))
//	}
//
//	fmt.Println("Staring json...")
//	response, _ := http.NewRequest("POST", url+"/tokens/isb?authtoken="+token, nil)

//	jsonData := map[string]string{"user": "teste"}
//	jsonValue, _ := json.Marshal(jsonData)
//	response, err = http.Post(url+"/tokens?authtoken="+token, "application/json", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Printf("Failed with error %s\n", err)
//	} else {
//		data, _ := ioutil.ReadAll(response.Body)
//		fmt.Println(string(data))
//	}

//}
