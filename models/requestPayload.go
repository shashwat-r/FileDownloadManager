package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestPayload struct {
	Type string `json:"type"`
	URLs []string `json:"urls"`
}

func (req RequestPayload) Get (r *http.Request) RequestPayload {
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("RPl", string(body))
	if err != nil {
		fmt.Println(err)
		return req
	}
	json.Unmarshal(body, &req)
	fmt.Println(req)
	return req
}


