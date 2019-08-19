package controller

import (
	"fmt"
	"github.com/shashwat-r/FileDownloadManager/download"
	"github.com/shashwat-r/FileDownloadManager/models"
	"net/http"
)


func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Download(w http.ResponseWriter, r *http.Request) {
	//req := models.GetRequestPayload(r)
	req := models.RequestPayload{}.Get(r)
	fmt.Printf("%s, %T\n", req.Type, req.Type)
	if req.Type == "serial" {
		SerialDownload(req.URLs)
	} else if req.Type == "concurrent" {
		ConcurrentDownload(req.URLs)
	} else {
		fmt.Println("Expected 'serial' or 'concurrent' in field 'type'")
	}
}

func SerialDownload(urls []string) {
	for _, url := range urls {
		err := download.Download(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ConcurrentDownload(urls []string) {
}

