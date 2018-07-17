package model

import "time"

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

//Jobs receives the Id and Name for the api jobs list.
type Jobs struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

//Execution is a nested sctruct that receives the results of the executions list.
type Execution struct {
	Paging struct {
		Count int `json:"count"`
		Total int `json:"total"`
	}
	Executions []struct {
		Id          int    `json:"id"`
		Status      string `json:"status"`
		Project     string `json:"project"`
		DateStarted struct {
			Date time.Time `json:"date"`
		} `json:"date-started"`
		//DateEnded struct {
		//	Date time.Time `json:"date"`
		//} `json:"date-ended"`
		Job struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Group string `json:"group"`
			//Description string `json:"description"`
		} `json:"job"`
	}
}
