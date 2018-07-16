package model

//ApiV is a struct which receive the rundeck api version
type ApiV struct {
	Version string `xml:"apiversion,attr"`
}

//Project receives a list of projects names
type Project struct {
	Count string   `xml:"count,attr"`
	Name  []string `xml:"project>name"`
}

//JVersion stores the Rundeck api version in json format
type JVersion struct {
	Vv int `json:"apiversion"`
}

//Projects receives the project name in json format
type Projects struct {
	Name string `json:"name"`
}

type Jobs struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Execution struct {
	Paging struct {
		Count int `json:"count"`
		Total int `json:"total"`
	}
	Executions []struct {
		Id      int    `json:"id"`
		Status  string `json:"status"`
		Project string `json:"project"`
	}
}