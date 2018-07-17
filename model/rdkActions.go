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
	"sync"
	"time"
)

var (
	Url,
	Token,
	Period,
	Max,
	Query,
	Action,
	Type,
	Name,
	ProjName string
	Syncer sync.WaitGroup
)

//Flags call the command-line flags arguments
func Flags() {
	flag.StringVar(&Url, "url", "http://localhost:4440", "the rundeck url")
	flag.StringVar(&Token, "token", "whSMNGB0UKuhItQnnzBn8qkCh4y3WFsy", "user auth token")

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
	flag.StringVar(&ProjName, "project", "", "The project name to narrow jobs list.")
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
	//	projectName := ListProjects(x, y, version)
	//fmt.Println(projectName)
	jsonOuts := []Jobs{}
	if ProjName == "" && Type != "proj" {
		fmt.Println("Error... You must specify a project name with -project \"NameOfMyproject\"!.")
		os.Exit(109)
	}
	//	for i := 0; i < len(projectName); i++ {
	//		req, err := http.NewRequest("GET", x+"/api/"+version+"/project/"+projectName[i].Name+"/jobs?authtoken="+y, nil)
	req, err := http.NewRequest("GET", x+"/api/"+version+"/project/"+ProjName+"/jobs?authtoken="+y, nil)
	Nerror(104, err, "[ListJobs] Fail when get reponse from jobs url. Error: ")
	req.Header.Set("Accept", "application/json")
	body := HttpClient(req)
	//Enable this print for debug proposes
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &jsonOuts)
	//	if len(jsonOuts) <= 0 {
	//		continue
	//	}
	//fmt.Println("Listing jobs in Project: ", string(projectName[i].Name))
	//}
	return (jsonOuts)
}

//ListExecutions receives two string url + token and return a list of executions narrow by flags.
func ListExecutions(x, y, z string) Execution {
	//Consult this nice curl converter on curl-to-Go: https://mholt.github.io/curl-to-go
	jsonOuts := Execution{}

	//params := strings.NewReader(`olderFilter=2w&max=0`)
	filter := ("olderFilter=" + Period + "&max=" + Max)
	if Query != "older" {
		filter = ("recentFilter=" + Period + "&max=" + Max)
	}
	if ProjName == "" && Type != "proj" {
		fmt.Println("Error... You must specify a project name with -project \"NameOfMyproject\"!.")
		os.Exit(109)
	}

	//	for i := 0; i < len(projectName); i++ {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	params := strings.NewReader(filter)
	req, err := http.NewRequest("POST", x+"/api/"+z+"/project/"+ProjName+"/executions", params)
	Nerror(105, err, "[ListOlderExecutions] Fail when get reponse from olderFilter url. Error: ")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Rundeck-Auth-Token", y)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	Nerror(106, err, "[ListOlderExecutions] Fail when execute request on url. Error: ")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Nerror(107, err, "[ListOlderExecutions] Fail when read the request url. Error: ")
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &jsonOuts)
	//	}
	//fmt.Printf("%+v\r\n", jsonOuts)
	//How to show only the id with this nested struct.
	//fmt.Println(jsonOuts.Executions[0].Id)
	return (jsonOuts)
}

//DeleteExecution receives ListExecutions and execute the delete for this execution
func DeleteExecution(id int, v string) {
	//version := strconv.Itoa(Version(Url))
	client := &http.Client{
		Timeout: time.Second * 900,
	}
	req, err := http.NewRequest("DELETE", Url+"/api/"+v+"/execution/"+strconv.Itoa(id), nil)
	Nerror(105, err, "[DeleteExecutions] Fail on delete request. Error: ")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Rundeck-Auth-Token", Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err = client.Do(req)
	Nerror(106, err, "[DeleteExecution] Fail when execute request on url. Error: ")
	fmt.Printf("Deleting Execution ID: [%d]\r\n", id)
	defer Syncer.Done()
	return
}

//Actions receives a flag string to run an action like: list or delete.
func Actions(x, y string) {
	ApiVersion := strconv.Itoa(Version(Url))
	List := ListExecutions(Url, Token, ApiVersion)
	if x == "list" {

		if y == "exec" {
			fmt.Println("Listing Executions...")

			//If Name are setted, list only executions from the ~project~job Name
			if Name != "" {
				for i := 0; i < len(List.Executions); i++ {
					//fmt.Println(list.Executions[i].Id)
					//if list.Executions[i].Project == Name {
					if List.Executions[i].Job.Name == Name {
						fmt.Printf("%+v\r\n", List.Executions[i])
					}
				}
				//Else list all the executions from all projects
			} else {
				for i := 0; i < len(List.Executions); i++ {
					//fmt.Println(list.Executions[i].Id)
					fmt.Printf("%+v\r\n", List.Executions[i])
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
			//			version := strconv.Itoa(Version(Url))
			projectName := ListProjects(Url, Token, ApiVersion)
			for i := 0; i < len(projectName); i++ {
				fmt.Println(projectName[i])
			}
		}
	}
	if x == "delete" {
		if y == "exec" {
			//If Name are setted, list only executions from the ~project~job Name
			if Name != "" {
				Syncer.Add(len(List.Executions))
				for i := 0; i < len(List.Executions); i++ {
					if List.Executions[i].Job.Name == Name {
						//fmt.Printf("%+v\r\n", List.Executions[i])
						go DeleteExecution(List.Executions[i].Id, ApiVersion)
					}
					Syncer.Wait()
				}
				//Else list all the executions from all projects
			} else {
				Syncer.Add(len(List.Executions))
				for i := 0; i < len(List.Executions); i++ {
					go DeleteExecution(List.Executions[i].Id, ApiVersion)
				}
				Syncer.Wait()
			}
		} else {
			fmt.Println("Sorry can't execute delete action on the resource: ", Type)
		}
	}
	return
}
