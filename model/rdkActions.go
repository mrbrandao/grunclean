package model

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Url, Token, Period, Max, Query, Action, Type, Name string
)

//Flags call the command-line flags arguments
func Flags() {
	flag.StringVar(&Url, "url", "https://localhost:4440", "the rundeck url")
	flag.StringVar(&Token, "token", "GKrfka6yPg145IQuvvXZXbU2GxU5fKzJ", "user auth token")

	/*Flags:
	  Max 20 > will narrow executions on max 20 executions.
	  Period 1d|1w|1m|1y > will narrow the executions by time range like: 1d = 1 day.See more at http://rundeck.org/docs/api/#execution-query
	  Query older|newer > will narrow the executions by time the Period value listing older or newer executions.
	*/

	flag.StringVar(&Period, "period", "1d", "The period of time range to narrow the executions result")
	flag.StringVar(&Max, "max", "20", "Maximum number of executions. If 0 will list all.")
	flag.StringVar(&Query, "query", "older", "Query executions by older or newer of the \"period\" flag.")
	flag.StringVar(&Action, "action", "list", "The Action to be used. Can be list or delete.")
	flag.StringVar(&Type, "type", "proj", "Which is the type you want list? Can be: \"proj|exec|job\".")
	flag.StringVar(&Name, "name", "", "Narrow querys by the name of the project, execution or job.")
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

//ListProjects receives url + token + version as a string and returns a slice with Projects Names
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
	//fmt.Println("Listing the projects found...")

	return (jsonOuts)
}

//ListJobs is a list of jobs found with ListProjects, it receiveis two strings url + token
func ListJobs(x, y string) []Jobs {
	version := strconv.Itoa(Version(x))
	projectName := ListProjects(x, y, version)
	//fmt.Println(projectName)
	jsonOuts := []Jobs{}
	for i := 0; i < len(projectName); i++ {
		req, err := http.NewRequest("GET", x+"/api/"+version+"/project/"+projectName[i].Name+"/jobs?authtoken="+y, nil)
		Nerror(104, err, "[ListJobs] Fail when get reponse from jobs url. Error: ")
		req.Header.Set("Accept", "application/json")
		body := HttpClient(req)
		//Enable this print for debug proposes
		//fmt.Println(string(body))
		err = json.Unmarshal(body, &jsonOuts)
		if len(jsonOuts) <= 0 {
			continue
		}
		//fmt.Println("Listing jobs in Project: ", string(projectName[i].Name))
	}
	return (jsonOuts)
}

//ListExecutions receives two string url + token and return a list of executions narrow by flags.
func ListExecutions(x, y string) Execution {
	//Consult this nice curl converter on curl-to-Go: https://mholt.github.io/curl-to-go
	version := strconv.Itoa(Version(x))
	projectName := ListProjects(x, y, version)
	//fmt.Println(projectName)
	jsonOuts := Execution{}
	//params := strings.NewReader(`olderFilter=2w&max=0`)
	filter := ("olderFilter=" + Period + "&max=" + Max)
	if Query != "older" {
		filter = ("recentFilter=" + Period + "&max=" + Max)
	}

	for i := 0; i < len(projectName); i++ {
		client := &http.Client{
			Timeout: time.Second * 5,
		}
		params := strings.NewReader(filter)
		req, err := http.NewRequest("POST", x+"/api/"+version+"/project/"+projectName[i].Name+"/executions", params)
		Nerror(105, err, "[ListOlderExecutions] Fail when get reponse from olderFilter url. Error: ")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Rundeck-Auth-Token", y)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		Nerror(106, err, "[ListOlderExecutions] Fail when execute request on url. Error: ")
		body, err := ioutil.ReadAll(resp.Body)
		Nerror(107, err, "[ListOlderExecutions] Fail when read the request url. Error: ")
		defer resp.Body.Close()
		//fmt.Println(string(body))
		err = json.Unmarshal(body, &jsonOuts)
	}
	//fmt.Printf("%+v\r\n", jsonOuts)
	//How to show only the id with this nested struct.
	//fmt.Println(jsonOuts.Executions[0].Id)
	return (jsonOuts)
}

//Actions receives a flag string to run an action like: list or delete.
func Actions(x, y string) {
	if x == "list" {

		if y == "exec" {
			fmt.Println("Listing Executions...")

			//If Name are setted list only executions from the project Name
			if Name != "" {
				list := ListExecutions(Url, Token)
				for i := 0; i < len(list.Executions); i++ {
					//fmt.Println(list.Executions[i].Id)
					if list.Executions[i].Project == Name {
						fmt.Printf("%+v\r\n", list.Executions[i])
					}
				}
				//Else list all the executions from all projects
			} else {
				list := ListExecutions(Url, Token)
				for i := 0; i < len(list.Executions); i++ {
					//fmt.Println(list.Executions[i].Id)
					fmt.Printf("%+v\r\n", list.Executions[i])
				}
			}

		} else if y == "job" {
			fmt.Println("Listing jobs: ")
			jobs := ListJobs(Url, Token)
			if Name != "" {
				for i := 0; i < len(jobs); i++ {
					if jobs[i].Name == Name {
						fmt.Printf("%+v\r\n", jobs[i])
					}
				}
			} else {
				for i := 0; i < len(jobs); i++ {
					fmt.Printf("%+v\r\n", jobs[i])
				}
			}
		} else if y == "proj" {
			fmt.Println("Listing all projects...")
			version := strconv.Itoa(Version(Url))
			projectName := ListProjects(Url, Token, version)
			for i := 0; i < len(projectName); i++ {
				fmt.Println(projectName[i])
			}
		}
	}
	return
}
