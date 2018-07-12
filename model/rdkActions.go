package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//Nerror treat the errors is more a reuse function to avoid excessive if err != nil on the code
func Nerror(i int, e error, s string) (error, string) {
	if e != nil {
		fmt.Println(s, e.Error())
		os.Exit(i)
	}
	return e, s
}

//HttpClient create the http client connection
func HttpClient(r *http.Request) []byte {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	client.Do(r)
	reply, err := client.Do(r)
	Nerror(101, err, "[HttpClient] Fail to request url with the Error: ")
	defer reply.Body.Close()
	body, _ := ioutil.ReadAll(reply.Body)
	return body
}

//Version returns the api version of Rundeck
func Version(x string) int {
	request, err := http.NewRequest("GET", x+"/api", nil)
	Nerror(102, err, "[Version] Fail on create the request to the rundeck /api url. Error: ")

	//Setting Header to response in Json Format
	request.Header.Set("Accept", "application/json")

	//Usefull to show type of something
	//fmt.Println(reflect.TypeOf(request))

	body := HttpClient(request)
	//fmt.Printf(string(body))
	jsonOut := JVersion{}
	err = json.Unmarshal(body, &jsonOut)
	return (jsonOut.Vv)
}

//ListProjects receives url + token + version in string and returns a list of Projects Names
func ListProjects(x, y, z string) []Projects {
	req, err := http.NewRequest("GET", x+"/api/"+z+"/projects?authtoken="+y, nil)
	Nerror(103, err, "[ListProjects] Fail when get reponse from projects url. Error: ")
	req.Header.Set("Accept", "application/json")
	body := HttpClient(req)

	//Making a slice from the struct also can be []struct check json array on json-to-go website
	jsonOuts := []Projects{}
	err = json.Unmarshal(body, &jsonOuts)

	//Example how to looping over this slice
	//fmt.Println(len(jsonOuts), jsonOuts[0].Name)
	//fmt.Println(reflect.ValueOf(jsonOuts).Kind())
	fmt.Println("Listing the projects found...")
	//for i := 0; i < len(jsonOuts); i++ {
	//	fmt.Println(jsonOuts[i].Name)
	//}

	return (jsonOuts)
}

//func ListJobs(x, y string) string {
//	for i := 0; i
//
//}
