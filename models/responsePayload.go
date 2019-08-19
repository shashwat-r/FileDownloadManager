package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ResponsePayload struct {
	ID           string `json:"id"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
	DownloadType string `json:"download_type"`
	Files        map[string]string `json:"files"`
}

var fileName = "downloads.json"

func (res ResponsePayload) Get() {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//var res ResponsePayload
	json.Unmarshal(byteValue, &res)
}

func (res ResponsePayload) Update(urlIdMap map[string]string) {
	res.Get()
	for k,v := range urlIdMap {
		res.Files[k] = v
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, resJson, 0644)
}
