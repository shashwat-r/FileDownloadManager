package main

import (
	"fmt"
	"github.com/shashwat-r/FileDownloadManager/controller"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hi")
	http.HandleFunc("/", controller.Router)
	log.Fatal(http.ListenAndServe(":8081", nil))
	//models.Read()
}