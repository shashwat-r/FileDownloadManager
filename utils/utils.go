package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func UUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func FileDownload(url string) (string, error) {
	fileId := UUID()
	//dir, err := filepath.Abs(filepath.Dir("main.go"))
	dir := "resources/tmp/"
	path := dir+fileId
	resp, err := http.Get(url)
	if err != nil {
		return fileId, err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return fileId, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return fileId, err
}

func ReadJsonBytesFromFile(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil
	}
	return jsonBytes
}

func WriteJsonBytesToFile(fileName string, bytes []byte) {
	_ = ioutil.WriteFile(fileName, bytes, 0644)
}
