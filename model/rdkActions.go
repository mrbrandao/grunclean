package model

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	Url, Token string
)

func Flags() {
	flag.StringVar(&Url, "url", "https://localhost:4440", "the rundeck url")
	flag.StringVar(&Token, "token", "GKrfka6yPg145IQuvvXZXbU2GxU5fKzJ", "user auth token")
	defer flag.Parse()
	return
}

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

//ListProjects receives url + token + version as a string and returns a slice of Projects Names
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

	return (jsonOuts)
}

//ListJobs is a list of jobs found with ListProjects it receiveis two strings url + token
func ListJobs(x, y string) string {
	version := strconv.Itoa(Version(Url))
	projectName := ListProjects(Url, Token, version)
	fmt.Println(projectName)
	for i := 0; i < len(projectName); i++ {
		fmt.Println("Listing jobs in Project: ", string(projectName[i].Name))
		req, err := http.NewRequest("GET", x+"/api/"+version+"/project/"+projectName[i].Name+"/jobs?authtoken="+y, nil)
		Nerror(104, err, "[ListJobs] Fail when get reponse from jobs url. Error: ")
		req.Header.Set("Accept", "application/json")
		body := HttpClient(req)
		fmt.Println(string(body))
	}
	return ("lalalala")

}
