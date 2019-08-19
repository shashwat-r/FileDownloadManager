package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Files struct {
	IdToURL, URLToId map[string]string
}

func Read() {
	//var files Files
	jsonFile, err := os.Open("download/files.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonFile)
	bytevalue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(bytevalue))
	var files Files
	json.Unmarshal(bytevalue, &files)
}