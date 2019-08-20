package models

import (
	"encoding/json"
	"github.com/shashwat-r/FileDownloadManager/utils"
	"strings"
	"time"
)

type ResponsePayload struct {
	ID           string `json:"id"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
	DownloadType string `json:"download_type"`
	Files        map[string]string `json:"files"`
}

func (res *ResponsePayload) Begin(downloadId, downloadType string) {
	res.ID = downloadId
	res.StartTime =time.Now().String()
	res.EndTime =time.Now().String()
	res.Status ="QUEUED"
	res.DownloadType = strings.ToUpper(downloadType)
	res.Files =make(map[string]string)
}

func (res *ResponsePayload) getFilePath() string {
	dir := "resources/info/"
	filePath := dir+res.ID+".json"
	return filePath
}

//func (res ResponsePayload) CreateFromFile(fileName string) {
//	jsonBytes := controller.ReadJsonBytesFromFile(fileName)
//	_ = json.Unmarshal(jsonBytes, &res)
//}

func (res *ResponsePayload) UpdateJson() {
	fileName := res.getFilePath()
	res.EndTime = time.Now().String()
	jsonBytes, _ := json.Marshal(res)
	utils.WriteJsonBytesToFile(fileName, jsonBytes)
}

func (res *ResponsePayload) UpdateFile(url, fileId string) {
	res.Files[url] = fileId
}

func (res *ResponsePayload) setEndStatus(status string) {
	res.Status = status
}

func (res *ResponsePayload) SetSuccessful() {
	res.setEndStatus("SUCCESSFUL")
}

func (res *ResponsePayload) SetFailed() {
	res.setEndStatus("FAILED")
}
