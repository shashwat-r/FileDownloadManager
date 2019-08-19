package download

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func Download(url string) error {
	//url := "https://getintopc.com/wp-content/uploads/2018/10/Nuance-Dragon-Professional-Individual-15-Free-Download.png"
	fileId := UUID()
	dir, err := filepath.Abs(filepath.Dir("main.go"))
	path := dir+"/tmp/"+fileId
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}