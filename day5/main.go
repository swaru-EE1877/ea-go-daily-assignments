package main

import (
	"encoding/json"
	"io/ioutil"
)

type Job struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type JobResult struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func main() {
	file, _ := ioutil.ReadFile("./jobs.json")

	var data []Job
	_ = json.Unmarshal(file, &data)

	chn := make(chan JobResult)
	for _, job := range data {
		go func(job Job) {
			chn <- execute(job)
		}(job)
	}

	var resultData []JobResult
	for _, _ = range data {
		resultData = append(resultData, <-chn)
	}
	result, _ := json.Marshal(resultData)
	ioutil.WriteFile("./result.json", result, 0644)
}

func execute(job Job) JobResult {
	return JobResult{Id: job.Id, Status: "SUCCESS"}
}
