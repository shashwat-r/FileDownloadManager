package controller

import (
	"encoding/json"
	"fmt"
	"github.com/shashwat-r/FileDownloadManager/models"
	"github.com/shashwat-r/FileDownloadManager/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Download(w http.ResponseWriter, r *http.Request) {
	//req := models.GetRequestPayload(r)
	req := models.RequestPayload{}.Get(r)
	fmt.Printf("%s, %T\n", req.Type, req.Type)
	if req.Type == "serial" {
		serialDownload(w, req.URLs)
	} else if req.Type == "concurrent" {
		concurrentDownload(w, req.URLs)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
  "internal_code": 4001,
  "message": "unknown type of download"
}`))
	}
}

func serialDownload(w http.ResponseWriter, urls []string) {
	downloadId := utils.UUID()
	status := models.ResponsePayload{}
	status.Begin(downloadId, "SERIAL")
	var count int
	for _, url := range urls {
		fileId, err := utils.FileDownload(url)
		if err != nil {
			fmt.Println(err)
		} else {
			count ++
		}
		status.UpdateFile(url, fileId)
	}
	if count == len(urls) {
		status.SetSuccessful()
	} else {
		status.SetFailed()
	}
	fmt.Println(status)
	type id struct {
		Id string
	}
	dId := id{Id:downloadId}
	fmt.Printf("%+v\n", dId)
	byteJson, err := json.Marshal(dId)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(byteJson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteJson)
}

func concurrentDownload(w http.ResponseWriter, urls []string) {
	downloadId := utils.UUID()
	status := models.ResponsePayload{}
	status.Begin(downloadId, "CONCURRENT")

	type id struct {
		Id string
	}
	dId := id{Id:downloadId}
	fmt.Printf("%+v\n", dId)
	byteJson, err := json.Marshal(dId)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(byteJson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteJson)

	urlChan := make(chan string)

	numThreads := 2
	for i:=0; i<numThreads; i++ {
		go parDownload(&status, urlChan)
	}
	go pushUrl(&status, urls, urlChan)
	//count = <- cntChan
	//if count == len(urls) {
	//	status.SetSuccessful()
	//} else {
	//	status.SetFailed()
	//}
}

func parDownload(status *models.ResponsePayload, cURL chan string) {
	for {
		url, ok := <-cURL
		if !ok {
			return
		}
		fileId, err := utils.FileDownload(url)
		if err != nil {
			fmt.Println(err)
		//} else {
		//	count := <-cCnt
		//	count++
		//	cCnt <- count
		}
		fmt.Println(url, fileId)
		status.UpdateFile(url, fileId)
		fmt.Println(status)
		status.SetSuccessful()
		status.UpdateJson()
	}
}

func pushUrl(status *models.ResponsePayload, urls []string, cURL chan string) {
	for _, url := range urls {
		cURL <- url
		fmt.Println("\n", url, "\n")
	}
	close(cURL)
	fmt.Println("closed curl")
	status.SetSuccessful()
	status.UpdateJson()
	fmt.Println(status.Status, status.Files)
}

func WriteJson(w http.ResponseWriter, id string) {
	filePath := "resources/info/"+id+".json"
	byteJson := utils.ReadJsonBytesFromFile(filePath)
	os.Stdout.Write(byteJson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if byteJson == nil {
		w.Write([]byte(`{
  "internal_code": 4002,
  "message": "unknown download ID"
}`))
	}
	w.Write(byteJson)
}

func ListFiles(w http.ResponseWriter) {
	files, err := ioutil.ReadDir("resources/tmp/")
	if err != nil {
		log.Fatal(err)
	}
	type file struct {
		Files []string
	}
	allFiles := "<pre>\n"
	for _, f := range files {
		n := f.Name()
		newString := "<a href=\""+n+"\">"+n+"</a>\n"
		allFiles += newString
	}
	allFiles += "</pre>"
	w.Write([]byte(allFiles))
}